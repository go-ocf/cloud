package service

import (
	"context"

	"github.com/plgd-dev/cloud/grpc-gateway/pb"
	"github.com/plgd-dev/cloud/pkg/log"
	kitNetGrpc "github.com/plgd-dev/cloud/pkg/net/grpc"
	"google.golang.org/grpc/codes"
)

func (r *RequestHandler) GetClientConfiguration(ctx context.Context, req *pb.ClientConfigurationRequest) (*pb.ClientConfigurationResponse, error) {
	ret, err := r.resourceDirectoryClient.GetClientConfiguration(ctx, req)
	if err != nil {
		return ret, log.LogAndReturnError(kitNetGrpc.ForwardErrorf(codes.Internal, "cannot get client configuration: %v", err))
	}
	return ret, err
}
