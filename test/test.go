package test

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/go-ocf/kit/codec/cbor"
	"github.com/go-ocf/kit/security/certManager"
	"go.uber.org/atomic"

	"github.com/go-ocf/kit/log"
	kitNetGrpc "github.com/go-ocf/kit/net/grpc"
	"github.com/go-ocf/kit/security"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/go-ocf/sdk/schema"
	"github.com/go-ocf/sdk/schema/cloud"

	coapgwService "github.com/go-ocf/cloud/coap-gateway/test"
	"github.com/go-ocf/cloud/grpc-gateway/pb"
	raService "github.com/go-ocf/cloud/resource-aggregate/test"
	rdService "github.com/go-ocf/cloud/resource-directory/test"
	"github.com/go-ocf/sdk/local/core"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	authService "github.com/go-ocf/cloud/authorization/test"
	grpcgwService "github.com/go-ocf/cloud/grpc-gateway/test"

	"github.com/go-ocf/cloud/authorization/provider"
)

var (
	TestDeviceName string

	TestDevsimResources        []schema.ResourceLink
	TestDevsimBackendResources []schema.ResourceLink
)

func init() {
	TestDeviceName = "devsim-" + MustGetHostname()
	TestDevsimResources = []schema.ResourceLink{
		{
			Href:          "/oic/p",
			ResourceTypes: []string{"oic.wk.p"},
			Interfaces:    []string{"oic.if.r", "oic.if.baseline"},
			Policy: schema.Policy{
				BitMask: 1,
			},
		},

		{
			Href:          "/oic/d",
			ResourceTypes: []string{"oic.d.cloudDevice", "oic.wk.d"},
			Interfaces:    []string{"oic.if.r", "oic.if.baseline"},
			Policy: schema.Policy{
				BitMask: 1,
			},
		},

		{
			Href:          "/oc/con",
			ResourceTypes: []string{"oic.wk.con"},
			Interfaces:    []string{"oic.if.rw", "oic.if.baseline"},
			Policy: schema.Policy{
				BitMask: 3,
			},
		},

		{
			Href:          "/light/1",
			ResourceTypes: []string{"core.light"},
			Interfaces:    []string{"oic.if.rw", "oic.if.baseline"},
			Policy: schema.Policy{
				BitMask: 3,
			},
		},

		{
			Href:          "/light/2",
			ResourceTypes: []string{"core.light"},
			Interfaces:    []string{"oic.if.rw", "oic.if.baseline"},
			Policy: schema.Policy{
				BitMask: 3,
			},
		},
	}

	TestDevsimBackendResources = []schema.ResourceLink{
		{
			Href:          cloud.StatusHref,
			ResourceTypes: cloud.StatusResourceTypes,
			Interfaces:    cloud.StatusInterfaces,
			Policy: schema.Policy{
				BitMask: 3,
			},
			Title: "Cloud device status",
		},
	}
}

func FindResourceLink(href string) schema.ResourceLink {
	for _, l := range TestDevsimResources {
		if l.Href == href {
			return l
		}
	}
	for _, l := range TestDevsimBackendResources {
		if l.Href == href {
			return l
		}
	}
	panic(fmt.Sprintf("resource %v: not found", href))
}

func ClearDB(ctx context.Context, t *testing.T) {
	var cmconfig certManager.Config
	err := envconfig.Process("DIAL", &cmconfig)
	assert.NoError(t, err)
	dialCertManager, err := certManager.NewCertManager(cmconfig)
	require.NoError(t, err)
	tlsConfig := dialCertManager.GetClientTLSConfig()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetTLSConfig(tlsConfig))
	require.NoError(t, err)
	dbs, err := client.ListDatabaseNames(ctx, bson.M{})
	if mongo.ErrNilDocument == err {
		return
	}
	require.NoError(t, err)
	for _, db := range dbs {
		if db == "admin" {
			continue
		}
		err = client.Database(db).Drop(ctx)
		require.NoError(t, err)
	}
	err = client.Disconnect(ctx)
	require.NoError(t, err)
}

func SetUp(ctx context.Context, t *testing.T) (TearDown func()) {
	testCfg.ClearDB(ctx, t)
	authShutdown := authService.SetUp(ctx, t)
	rdShutdown := rdService.SetUp(ctx, t)
	raShutdown := raService.SetUp(ctx, t)
	gwShutdown := coapgwService.SetUp(ctx, t)
	grpcShutdown := grpcgwService.SetUp(t, grpcCfg)

	return func() {
		grpcShutdown()
		gwShutdown()
		raShutdown()
		rdShutdown()
		authShutdown()
	}
}

func OnboardDevSim(ctx context.Context, t *testing.T, c pb.GrpcGatewayClient, deviceID string, gwHost string, expectedResources []schema.ResourceLink) func() {
	log.Setup(log.Config{Debug: false})
	client := core.NewClient()
	dev, links, err := client.GetDevice(ctx, deviceID)
	require.NoError(t, err)
	devLink, ok := links.GetResourceLink("/oic/d")
	require.True(t, ok)
	patchedLinks := make(schema.ResourceLinks, 0, len(links))
	for _, l := range links {
		if len(l.Endpoints) == 0 {
			l.Endpoints = devLink.Endpoints
		}
		patchedLinks = append(patchedLinks, l)
	}
	link, ok := patchedLinks.GetResourceLink("/CoapCloudConfResURI")
	require.True(t, ok)

	err = dev.UpdateResource(ctx, link, cloud.ConfigurationUpdateRequest{
		AuthorizationProvider: "test",
		URL:                   "coap+tcp://" + gwHost,
		AuthorizationCode:     "authCode",
		CloudID:               "sid",
	}, nil)
	require.NoError(t, err)

	waitForDevice(ctx, t, c, deviceID, expectedResources)

	return func() {
		err = dev.UpdateResource(ctx, link, cloud.ConfigurationUpdateRequest{
			AuthorizationProvider: "",
			URL:                   "",
			AuthorizationCode:     "",
		}, nil)
		require.NoError(t, err)
		dev.Close(ctx)
	}
}

func waitForDevice(ctx context.Context, t *testing.T, c pb.GrpcGatewayClient, deviceID string, expectedResources []schema.ResourceLink) {
	ctx = kitNetGrpc.CtxWithToken(ctx, provider.UserToken)
	client, err := c.SubscribeForEvents(ctx)
	require.NoError(t, err)

	err = client.Send(&pb.SubscribeForEvents{
		Token: "testToken",
		FilterBy: &pb.SubscribeForEvents_DevicesEvent{
			DevicesEvent: &pb.SubscribeForEvents_DevicesEventFilter{
				FilterEvents: []pb.SubscribeForEvents_DevicesEventFilter_Event{
					pb.SubscribeForEvents_DevicesEventFilter_ONLINE,
				},
			},
		},
	})
	require.NoError(t, err)
	ev, err := client.Recv()
	require.NoError(t, err)
	expectedEvent := &pb.Event{
		SubscriptionId: ev.SubscriptionId,
		Type: &pb.Event_OperationProcessed_{
			OperationProcessed: &pb.Event_OperationProcessed{
				Token: "testToken",
				ErrorStatus: &pb.Event_OperationProcessed_ErrorStatus{
					Code: pb.Event_OperationProcessed_ErrorStatus_OK,
				},
			},
		},
	}
	require.Equal(t, expectedEvent, ev)

	ev, err = client.Recv()
	require.NoError(t, err)
	expectedEvent = &pb.Event{
		SubscriptionId: ev.SubscriptionId,
		Type: &pb.Event_DeviceOnline_{
			DeviceOnline: &pb.Event_DeviceOnline{
				DeviceId: deviceID,
			},
		},
	}
	require.Equal(t, expectedEvent, ev)

	err = client.Send(&pb.SubscribeForEvents{
		Token: "testToken",
		FilterBy: &pb.SubscribeForEvents_DeviceEvent{
			DeviceEvent: &pb.SubscribeForEvents_DeviceEventFilter{
				DeviceId: deviceID,
				FilterEvents: []pb.SubscribeForEvents_DeviceEventFilter_Event{
					pb.SubscribeForEvents_DeviceEventFilter_RESOURCE_PUBLISHED,
				},
			},
		},
	})
	require.NoError(t, err)
	ev, err = client.Recv()
	require.NoError(t, err)
	expectedEvent = &pb.Event{
		Type: &pb.Event_OperationProcessed_{
			OperationProcessed: &pb.Event_OperationProcessed{
				Token: "testToken",
				ErrorStatus: &pb.Event_OperationProcessed_ErrorStatus{
					Code: pb.Event_OperationProcessed_ErrorStatus_OK,
				},
			},
		},
	}
	expectedEvent.SubscriptionId = ev.SubscriptionId
	require.Equal(t, expectedEvent, ev)
	subOnPublishedId := ev.SubscriptionId

	expectedEvents := ResourceLinksToExpectedPublishEvents(deviceID, expectedResources)

	for len(expectedEvents) > 0 {
		ev, err = client.Recv()
		require.NoError(t, err)
		ev.SubscriptionId = ""
		key := ev.GetResourcePublished().GetLink().GetDeviceId() + ev.GetResourcePublished().GetLink().GetHref()
		expectedEvents[key].GetResourcePublished().GetLink().InstanceId = ev.GetResourcePublished().GetLink().GetInstanceId()

		require.Equal(t, expectedEvents[key], ev)
		delete(expectedEvents, key)
	}

	err = client.Send(&pb.SubscribeForEvents{
		Token: "testToken",
		FilterBy: &pb.SubscribeForEvents_CancelSubscription_{
			CancelSubscription: &pb.SubscribeForEvents_CancelSubscription{
				SubscriptionId: subOnPublishedId,
			},
		},
	})
	require.NoError(t, err)

	ev, err = client.Recv()
	require.NoError(t, err)
	expectedEvent = &pb.Event{
		SubscriptionId: ev.SubscriptionId,
		Type: &pb.Event_SubscriptionCanceled_{
			SubscriptionCanceled: &pb.Event_SubscriptionCanceled{},
		},
	}
	require.Equal(t, expectedEvent, ev)

	ev, err = client.Recv()
	require.NoError(t, err)
	expectedEvent = &pb.Event{
		SubscriptionId: ev.SubscriptionId,
		Type: &pb.Event_OperationProcessed_{
			OperationProcessed: &pb.Event_OperationProcessed{
				Token: "testToken",
				ErrorStatus: &pb.Event_OperationProcessed_ErrorStatus{
					Code: pb.Event_OperationProcessed_ErrorStatus_OK,
				},
			},
		},
	}
	require.Equal(t, expectedEvent, ev)

	err = client.CloseSend()
	require.NoError(t, err)
}

func GetRootCertificatePool(t *testing.T) *x509.CertPool {
	pool := security.NewDefaultCertPool(nil)
	dat, err := ioutil.ReadFile(os.Getenv("LISTEN_ACME_CA_POOL"))
	require.NoError(t, err)
	ok := pool.AppendCertsFromPEM(dat)
	require.True(t, ok)
	return pool
}

func GetRootCertificateAuthorities(t *testing.T) []*x509.Certificate {
	dat, err := ioutil.ReadFile(os.Getenv("LISTEN_ACME_CA_POOL"))
	require.NoError(t, err)
	r := make([]*x509.Certificate, 0, 4)
	for {
		block, rest := pem.Decode(dat)
		require.NotNil(t, block)
		certs, err := x509.ParseCertificates(block.Bytes)
		require.NoError(t, err)
		r = append(r, certs...)
		if len(rest) == 0 {
			break
		}
	}

	return r
}

func MustGetHostname() string {
	n, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return n
}

func MustFindDeviceByName(name string) (deviceID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	deviceID, err := FindDeviceByName(ctx, name)
	if err != nil {
		panic(err)
	}
	return deviceID
}

type findDeviceIDByNameHandler struct {
	id     atomic.Value
	name   string
	cancel context.CancelFunc
}

func (h *findDeviceIDByNameHandler) Handle(ctx context.Context, device *core.Device, deviceLinks schema.ResourceLinks) {
	defer device.Close(ctx)
	l, ok := deviceLinks.GetResourceLink("/oic/d")
	if !ok {
		return
	}
	var d schema.Device
	err := device.GetResource(ctx, l, &d)
	if err != nil {
		return
	}
	if d.Name == h.name {
		h.id.Store(d.ID)
		h.cancel()
	}
}

func (h *findDeviceIDByNameHandler) Error(err error) {}

func FindDeviceByName(ctx context.Context, name string) (deviceID string, _ error) {
	client := core.NewClient()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	h := findDeviceIDByNameHandler{
		name:   name,
		cancel: cancel,
	}

	err := client.GetDevices(ctx, &h)
	if err != nil {
		return "", fmt.Errorf("could not find the device named %s: %w", name, err)
	}
	id, ok := h.id.Load().(string)
	if !ok || id == "" {
		return "", fmt.Errorf("could not find the device named %s: not found", name)
	}
	return id, nil
}

func DecodeCbor(t *testing.T, data []byte) interface{} {
	var v interface{}
	err := cbor.Decode(data, &v)
	require.NoError(t, err)
	return v
}

func EncodeToCbor(t *testing.T, v interface{}) []byte {
	d, err := cbor.Encode(v)
	require.NoError(t, err)
	return d
}

func ResourceLinkToPublishEvent(deviceID string, instanceID int64, l schema.ResourceLink) *pb.Event {
	link := pb.SchemaResourceLinkToProto(l)
	link.DeviceId = deviceID
	link.InstanceId = instanceID
	return &pb.Event{
		Type: &pb.Event_ResourcePublished_{
			ResourcePublished: &pb.Event_ResourcePublished{
				Link: &link,
			},
		},
	}
}

func ResourceLinksToExpectedPublishEvents(deviceID string, links []schema.ResourceLink) map[string]*pb.Event {
	e := make(map[string]*pb.Event)
	for _, l := range links {
		e[deviceID+l.Href] = ResourceLinkToPublishEvent(deviceID, 0, l)
	}
	return e
}

func GetAllBackendResourceLinks() []schema.ResourceLink {
	return append(TestDevsimResources, TestDevsimBackendResources...)
}

func ResourceLinksToPb(deviceID string, s []schema.ResourceLink) []pb.ResourceLink {
	r := make([]pb.ResourceLink, 0, len(s))
	for _, l := range s {
		l.DeviceID = deviceID
		r = append(r, pb.SchemaResourceLinkToProto(l))
	}
	return r
}

type SortResourcesByHref []pb.ResourceLink

func (a SortResourcesByHref) Len() int      { return len(a) }
func (a SortResourcesByHref) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortResourcesByHref) Less(i, j int) bool {
	return a[i].Href < a[j].Href
}

func SortResources(s []pb.ResourceLink) []pb.ResourceLink {
	v := SortResourcesByHref(s)
	sort.Sort(v)
	return v
}
