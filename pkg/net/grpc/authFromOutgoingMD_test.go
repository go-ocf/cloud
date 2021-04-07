package grpc_test

import (
	"context"
	"testing"

	"github.com/plgd-dev/cloud/pkg/net/grpc"
	"github.com/stretchr/testify/require"
)

func TestAuthFromOutgoingMD(t *testing.T) {
	token := "token"
	got, err := grpc.TokenFromOutgoingMD(grpc.CtxWithToken(context.Background(), token))
	require.NoError(t, err)
	require.Equal(t, token, got)
}

func TestOwnerFromOutgoingMD(t *testing.T) {
	token := "token"
	got, err := grpc.OwnerFromOutgoingMD(grpc.CtxWithOwner(context.Background(), token))
	require.NoError(t, err)
	require.Equal(t, token, got)
}

func TestOwnerFromMD(t *testing.T) {
	token := "token"
	got, err := grpc.OwnerFromMD(grpc.CtxWithIncomingOwner(context.Background(), token))
	require.NoError(t, err)
	require.Equal(t, token, got)
}

func TestAuthIDFromMD(t *testing.T) {
	token := "token"
	got, err := grpc.TokenFromMD(grpc.CtxWithIncomingToken(context.Background(), token))
	require.NoError(t, err)
	require.Equal(t, token, got)
}
