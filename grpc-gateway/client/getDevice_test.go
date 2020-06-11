package client_test

import (
	"context"
	"crypto/x509"
	"testing"
	"time"

	kitNetGrpc "github.com/go-ocf/kit/net/grpc"

	authTest "github.com/go-ocf/cloud/authorization/provider"
	"github.com/go-ocf/cloud/grpc-gateway/client"
	"github.com/go-ocf/cloud/grpc-gateway/pb"
	test "github.com/go-ocf/cloud/test"
	testCfg "github.com/go-ocf/cloud/test/config"
	"github.com/stretchr/testify/require"
)

const TestTimeout = time.Second * 20
const DeviceSimulatorIdNotFound = "00000000-0000-0000-0000-000000000111"

type testApplication struct {
	cas []*x509.Certificate
}

func (a *testApplication) GetRootCertificateAuthorities() ([]*x509.Certificate, error) {
	return a.cas, nil
}

func NewTestDeviceSimulator(deviceID, deviceName string) client.DeviceDetails {
	return client.DeviceDetails{
		ID: deviceID,
		Device: pb.Device{
			Id:       deviceID,
			Name:     deviceName,
			IsOnline: true,
		},
		Resources: test.SortResources(test.ResourceLinksToPb(deviceID, test.GetAllBackendResourceLinks())),
	}
}

func TestClient_GetDevice(t *testing.T) {
	deviceID := test.MustFindDeviceByName(test.TestDeviceName)
	type args struct {
		token    string
		deviceID string
	}
	tests := []struct {
		name    string
		args    args
		want    client.DeviceDetails
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				token:    authTest.UserToken,
				deviceID: deviceID,
			},
			want: NewTestDeviceSimulator(deviceID, test.TestDeviceName),
		},
		{
			name: "not-found",
			args: args{
				token:    authTest.UserToken,
				deviceID: "not-found",
			},
			wantErr: true,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), TestTimeout)
	defer cancel()
	ctx = kitNetGrpc.CtxWithToken(ctx, authTest.UserToken)

	tearDown := test.SetUp(ctx, t)
	defer tearDown()

	c := NewTestClient(t)
	defer c.Close(context.Background())

	shutdownDevSim := test.OnboardDevSim(ctx, t, c.GrpcGatewayClient(), deviceID, testCfg.GW_HOST, test.GetAllBackendResourceLinks())
	defer shutdownDevSim()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()
			got, err := c.GetDevice(ctx, tt.args.deviceID)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			got.Resources = test.SortResources(got.Resources)
			require.Equal(t, tt.want, got)
		})
	}
}
