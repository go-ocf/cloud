package mongodb

import (
	"context"
	"testing"

	"github.com/kelseyhightower/envconfig"
	"github.com/plgd-dev/cloud/resource-aggregate/cqrs/eventstore/test"
	"github.com/plgd-dev/kit/security/certManager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestEventStore(t *testing.T) {
	var config certManager.Config
	err := envconfig.Process("DIAL", &config)
	assert.NoError(t, err)

	dialCertManager, err := certManager.NewCertManager(config)
	require.NoError(t, err)

	ctx := context.Background()
	tlsConfig := dialCertManager.GetClientTLSConfig()

	store, err := NewEventStore(
		Config{
			URI: "mongodb://localhost:27017",
		},
		func(f func()) error { go f(); return nil },
		WithMarshaler(bson.Marshal),
		WithUnmarshaler(bson.Unmarshal),
		WithTLS(tlsConfig),
	)
	assert.NoError(t, err)
	assert.NotNil(t, store)
	defer store.Close(ctx)
	defer func() {
		t.Log("clearing db")
		err := store.Clear(ctx)
		require.NoError(t, err)
	}()

	t.Log("event store with default namespace")
	test.AcceptanceTest(t, ctx, store)
}
