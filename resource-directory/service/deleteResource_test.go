package service_test

import (
	"context"
	"crypto/tls"
	"testing"

	"github.com/plgd-dev/cloud/authorization/provider"
	"github.com/plgd-dev/cloud/grpc-gateway/pb"
	extCodes "github.com/plgd-dev/cloud/grpc-gateway/pb/codes"
	"github.com/plgd-dev/cloud/test"
	testCfg "github.com/plgd-dev/cloud/test/config"
	"github.com/plgd-dev/kit/codec/cbor"
	"github.com/plgd-dev/kit/log"
	kitNetGrpc "github.com/plgd-dev/kit/net/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func TestRequestHandler_DeleteResource(t *testing.T) {
	deviceID := test.MustFindDeviceByName(test.TestDeviceName)
	type args struct {
		req pb.DeleteResourceRequest
	}
	tests := []struct {
		name            string
		args            args
		want            map[string]interface{}
		wantContentType string
		wantErr         bool
		wantErrCode     codes.Code
	}{
		{
			name: "/light/2 - MethodNotAllowed",
			args: args{
				req: pb.DeleteResourceRequest{
					ResourceId: &pb.ResourceId{
						DeviceId: deviceID,
						Href:     "/light/2",
					},
				},
			},
			wantErr:     true,
			wantErrCode: extCodes.MethodNotAllowed,
		},
		{
			name: "invalid Href",
			args: args{
				req: pb.DeleteResourceRequest{
					ResourceId: &pb.ResourceId{
						DeviceId: deviceID,
						Href:     "/unknown",
					},
				},
			},
			wantErr:     true,
			wantErrCode: codes.NotFound,
		},

		{
			name: "/oic/d - PermissionDenied",
			args: args{
				req: pb.DeleteResourceRequest{
					ResourceId: &pb.ResourceId{
						DeviceId: deviceID,
						Href:     "/oic/d",
					},
				},
			},
			wantErr:     true,
			wantErrCode: codes.PermissionDenied,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), TEST_TIMEOUT)
	defer cancel()
	ctx = kitNetGrpc.CtxWithToken(ctx, provider.UserToken)

	tearDown := test.SetUp(ctx, t)
	defer tearDown()

	conn, err := grpc.Dial(testCfg.GRPC_HOST, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		RootCAs: test.GetRootCertificatePool(t),
	})))
	require.NoError(t, err)
	c := pb.NewGrpcGatewayClient(conn)

	log.Setup(log.Config{Debug: true})
	deviceID, shutdownDevSim := test.OnboardDevSim(ctx, t, c, deviceID, testCfg.GW_HOST, test.GetAllBackendResourceLinks())
	defer shutdownDevSim()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.DeleteResource(ctx, &tt.args.req)
			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, tt.wantErrCode, status.Convert(err).Code())
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.wantContentType, got.GetContent().GetContentType())
				var d map[string]interface{}
				err := cbor.Decode(got.GetContent().GetData(), &d)
				require.NoError(t, err)
				delete(d, "piid")
				assert.Equal(t, tt.want, d)
			}
		})
	}
}
