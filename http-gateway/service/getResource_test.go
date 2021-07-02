package service_test

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/plgd-dev/cloud/grpc-gateway/pb"
	httpgwTest "github.com/plgd-dev/cloud/http-gateway/test"
	testHttp "github.com/plgd-dev/cloud/http-gateway/test"
	"github.com/plgd-dev/cloud/http-gateway/uri"
	kitNetGrpc "github.com/plgd-dev/cloud/pkg/net/grpc"
	"github.com/plgd-dev/cloud/resource-aggregate/commands"
	"github.com/plgd-dev/cloud/resource-aggregate/events"
	"github.com/plgd-dev/cloud/test"
	testCfg "github.com/plgd-dev/cloud/test/config"
	oauthTest "github.com/plgd-dev/cloud/test/oauth-server/test"
	"github.com/plgd-dev/go-coap/v2/message"
)

func NewBool(v bool) *bool {
	return &v
}

func cmpResourceChanged(t *testing.T, want, got *events.ResourceChanged) {
	dataWant := want.GetContent().GetData()
	datagot := got.GetContent().GetData()
	want.Content.Data = nil
	got.Content.Data = nil
	test.CheckProtobufs(t, want, got, test.RequireToCheckFunc(require.Equal))
	w := test.DecodeCbor(t, dataWant)
	g := test.DecodeCbor(t, datagot)
	require.Equal(t, w, g)
}

func TestRequestHandler_GetResource(t *testing.T) {
	deviceID := test.MustFindDeviceByName(test.TestDeviceName)
	type args struct {
		deviceID          string
		resourceHref      string
		shadow            *bool
		resourceInterface string
		accept            string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    *events.ResourceChanged
	}{
		{
			name: "jsonpb: get from resource shadow",
			args: args{
				deviceID:     deviceID,
				resourceHref: "/light/1",
				accept:       uri.ApplicationJsonPBContentType,
			},
			want: &events.ResourceChanged{
				ResourceId: &commands.ResourceId{
					DeviceId: deviceID,
					Href:     "/light/1",
				},
				Content: &commands.Content{
					CoapContentFormat: int32(message.AppOcfCbor),
					ContentType:       message.AppOcfCbor.String(),
					Data: test.EncodeToCbor(t, map[string]interface{}{
						"state": false,
						"power": uint64(0),
						"name":  "Light",
						"if":    []interface{}{"oic.if.rw", "oic.if.baseline"},
						"rt":    []interface{}{"core.light"},
					}),
				},
				Status: commands.Status_OK,
			},
		},
		{
			name: "jsonpb: get from device with interface",
			args: args{
				deviceID:          deviceID,
				resourceHref:      "/light/1",
				resourceInterface: "oic.if.baseline",
				accept:            uri.ApplicationJsonPBContentType,
			},
			want: &events.ResourceChanged{
				ResourceId: &commands.ResourceId{
					DeviceId: deviceID,
					Href:     "/light/1",
				},
				Content: &commands.Content{
					CoapContentFormat: int32(message.AppOcfCbor),
					ContentType:       message.AppOcfCbor.String(),
					Data: test.EncodeToCbor(t, map[string]interface{}{
						"state": false,
						"power": uint64(0),
						"name":  "Light",
						"if":    []interface{}{"oic.if.rw", "oic.if.baseline"},
						"rt":    []interface{}{"core.light"},
					}),
				},
				Status: commands.Status_OK,
			},
		},
		{
			name: "jsonpb: get from device with disabled shadow",
			args: args{
				deviceID:     deviceID,
				resourceHref: "/light/1",
				shadow:       NewBool(false),
				accept:       uri.ApplicationJsonPBContentType,
			},
			want: &events.ResourceChanged{
				ResourceId: &commands.ResourceId{
					DeviceId: deviceID,
					Href:     "/light/1",
				},
				Content: &commands.Content{
					CoapContentFormat: int32(message.AppOcfCbor),
					ContentType:       message.AppOcfCbor.String(),
					Data: test.EncodeToCbor(t, map[string]interface{}{
						"state": false,
						"power": uint64(0),
						"name":  "Light",
					}),
				},
				Status: commands.Status_OK,
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), testCfg.TEST_TIMEOUT)
	defer cancel()

	tearDown := test.SetUp(ctx, t)
	defer tearDown()

	shutdownHttp := testHttp.SetUp(t)
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
			req := httpgwTest.NewRequest(http.MethodGet, uri.AliasDeviceResource, nil).DeviceId(tt.args.deviceID).ResourceHref(tt.args.resourceHref).ResourceInterface(tt.args.resourceInterface).AuthToken(token).Accept(tt.args.accept)
			if tt.args.shadow != nil {
				req.Shadow(*tt.args.shadow)
			}
			request := req.Build()
			trans := http.DefaultTransport.(*http.Transport).Clone()
			trans.TLSClientConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
			c := http.Client{
				Transport: trans,
			}
			resp, err := c.Do(request)
			require.NoError(t, err)
			defer resp.Body.Close()

			marshaler := runtime.JSONPb{}
			decoder := marshaler.NewDecoder(resp.Body)
			values := make([]*events.ResourceChanged, 0, 1)
			for {
				var value events.ResourceChanged
				err = Unmarshal(resp.StatusCode, decoder, &value)
				if err == io.EOF {
					break
				}
				if tt.wantErr {
					require.Error(t, err)
					return
				}
				require.NoError(t, err)
				value.AuditContext = nil
				value.EventMetadata = nil
				values = append(values, &value)
			}
			cmpResourceChanged(t, tt.want, values[0])
		})
	}
}
