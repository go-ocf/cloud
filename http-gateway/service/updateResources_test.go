package service_test

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/plgd-dev/cloud/grpc-gateway/pb"
	"github.com/plgd-dev/cloud/http-gateway/service"
	kitNetGrpc "github.com/plgd-dev/cloud/pkg/net/grpc"
	"github.com/plgd-dev/cloud/resource-aggregate/commands"
	"github.com/plgd-dev/cloud/resource-aggregate/events"
	"github.com/plgd-dev/cloud/test"
	"github.com/plgd-dev/cloud/test/config"
	testCfg "github.com/plgd-dev/cloud/test/config"
	oauthTest "github.com/plgd-dev/cloud/test/oauth-server/test"
	"github.com/plgd-dev/go-coap/v2/message"
)

func updateResource(ctx context.Context, req *pb.UpdateResourceRequest, token string) (*events.ResourceUpdated, error) {
	var m jsonpb.Marshaler
	data, err := m.MarshalToString(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("https://%v/api/v1/devices/%v", config.HTTP_GW_HOST, req.ResourceId), bytes.NewReader([]byte(data)))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("bearer %s", token))
	trans := http.DefaultTransport.(*http.Transport).Clone()
	trans.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	c := http.Client{
		Transport: trans,
	}
	resp, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	marshaler := runtime.JSONPb{}
	decoder := marshaler.NewDecoder(resp.Body)

	var got events.ResourceUpdated
	err = service.Unmarshal(resp.StatusCode, decoder, &got)
	if err != nil {
		return nil, err
	}
	return &got, nil
}

func TestRequestHandler_UpdateResourcesValues(t *testing.T) {
	deviceID := test.MustFindDeviceByName(test.TestDeviceName)
	type args struct {
		req pb.UpdateResourceRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *events.ResourceUpdated
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				req: pb.UpdateResourceRequest{
					ResourceId: commands.NewResourceID(deviceID, "/light/1").ToString(),
					Content: &pb.Content{
						ContentType: message.AppOcfCbor.String(),
						Data: test.EncodeToCbor(t, map[string]interface{}{
							"power": 1,
						}),
					},
				},
			},
			want: &events.ResourceUpdated{
				ResourceId: &commands.ResourceId{
					DeviceId: deviceID,
					Href:     "/light/1",
				},
				Content: &commands.Content{
					CoapContentFormat: -1,
				},
				Status: commands.Status_OK,
			},
		},
		{
			name: "valid with interface",
			args: args{
				req: pb.UpdateResourceRequest{
					ResourceInterface: "oic.if.baseline",
					ResourceId:        commands.NewResourceID(deviceID, "/light/1").ToString(),
					Content: &pb.Content{
						ContentType: message.AppOcfCbor.String(),
						Data: test.EncodeToCbor(t, map[string]interface{}{
							"power": 2,
						}),
					},
				},
			},
			want: &events.ResourceUpdated{
				ResourceId: &commands.ResourceId{
					DeviceId: deviceID,
					Href:     "/light/1",
				},
				Content: &commands.Content{
					CoapContentFormat: -1,
				},
				Status: commands.Status_OK,
			},
		},
		{
			name: "revert update",
			args: args{
				req: pb.UpdateResourceRequest{
					ResourceInterface: "oic.if.baseline",
					ResourceId:        commands.NewResourceID(deviceID, "/light/1").ToString(),
					Content: &pb.Content{
						ContentType: message.AppOcfCbor.String(),
						Data: test.EncodeToCbor(t, map[string]interface{}{
							"power": 0,
						}),
					},
				},
			},
			want: &events.ResourceUpdated{
				ResourceId: commands.NewResourceID(deviceID, "/light/1"),
				Content: &commands.Content{
					CoapContentFormat: -1,
				},
				Status: commands.Status_OK,
			},
		},
		{
			name: "update RO-resource",
			args: args{
				req: pb.UpdateResourceRequest{
					ResourceId: commands.NewResourceID(deviceID, "/oic/d").ToString(),
					Content: &pb.Content{
						ContentType: message.AppOcfCbor.String(),
						Data: test.EncodeToCbor(t, map[string]interface{}{
							"di": "abc",
						}),
					},
				},
			},
			want: &events.ResourceUpdated{
				ResourceId: commands.NewResourceID(deviceID, "/oic/d"),
				Content: &commands.Content{
					CoapContentFormat: -1,
				},
				Status: commands.Status_FORBIDDEN,
			},
		},
		{
			name: "invalid Href",
			args: args{
				req: pb.UpdateResourceRequest{
					ResourceId: commands.NewResourceID(deviceID, "/unknown").ToString(),
				},
			},
			wantErr: true,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), testCfg.TEST_TIMEOUT)
	defer cancel()

	tearDown := test.SetUp(ctx, t)
	defer tearDown()

	shutdownHttp := New(t, MakeConfig(t))
	defer shutdownHttp()

	token := oauthTest.GetServiceToken(t)
	ctx = kitNetGrpc.CtxWithToken(ctx, token)

	conn, err := grpc.Dial(testCfg.GRPC_HOST, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		RootCAs: test.GetRootCertificatePool(t),
	})))
	require.NoError(t, err)
	c := pb.NewGrpcGatewayClient(conn)

	deviceID, shutdownDevSim := test.OnboardDevSim(ctx, t, c, deviceID, testCfg.GW_HOST, test.GetAllBackendResourceLinks())
	defer shutdownDevSim()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updateResource(ctx, &tt.args.req, token)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			got.AuditContext = nil
			got.EventMetadata = nil
			test.CheckProtobufs(t, tt.want, &got, test.RequireToCheckFunc(require.Equal))
		})
	}
}