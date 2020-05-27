package service

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/go-ocf/cloud/coap-gateway/coapconv"
	cqrsRA "github.com/go-ocf/cloud/resource-aggregate/cqrs"
	raEvents "github.com/go-ocf/cloud/resource-aggregate/cqrs/events"
	pbCQRS "github.com/go-ocf/cloud/resource-aggregate/pb"
	pbRA "github.com/go-ocf/cloud/resource-aggregate/pb"
	coapCodes "github.com/go-ocf/go-coap/v2/message/codes"
	"github.com/go-ocf/go-coap/v2/tcp"
	"github.com/go-ocf/go-coap/v2/tcp/message/pool"
	"github.com/go-ocf/kit/log"
	kitNetGrpc "github.com/go-ocf/kit/net/grpc"
	"github.com/go-ocf/sdk/schema/cloud"
)

type observedResource struct {
	res         *pbRA.Resource
	observation *tcp.Observation
}

type authCtx struct {
	pbCQRS.AuthorizationContext
	AccessToken string
}

//Client a setup of connection
type Client struct {
	server   *Server
	coapConn *tcp.ClientConn
	isClosed int32

	observedResources     map[string]map[int64]observedResource // [deviceID][instanceID]
	observedResourcesLock sync.Mutex
	authCtx               authCtx
	authContextLock       sync.Mutex
}

//newClient create and initialize client
func newClient(server *Server, client *tcp.ClientConn) *Client {
	return &Client{
		server:            server,
		coapConn:          client,
		observedResources: make(map[string]map[int64]observedResource),
	}
}

func (client *Client) remoteAddrString() string {
	return client.coapConn.RemoteAddr().String()
}

func (client *Client) observeResource(ctx context.Context, res *pbRA.Resource, allowDuplicit bool) (err error) {
	log.Debugf("DeviceId: %v, ResourceId: %v: observe resource", res.DeviceId, res.Id)

	client.observedResourcesLock.Lock()
	defer client.observedResourcesLock.Unlock()
	if _, ok := client.observedResources[res.DeviceId]; !ok {
		client.observedResources[res.DeviceId] = make(map[int64]observedResource)
	}
	if _, ok := client.observedResources[res.DeviceId][res.InstanceId]; ok {
		if allowDuplicit {
			return nil
		}
		return fmt.Errorf("resource is already already published")
	}
	return client.addObservedResourceLocked(ctx, res)
}

func (client *Client) getResourceContent(ctx context.Context, obsRes *pbRA.Resource) {
	resp, err := client.coapConn.Get(ctx, obsRes.Href)
	if err != nil {
		log.Errorf("DeviceId: %v, ResourceId: %v: cannot get resource content: %v", obsRes.DeviceId, obsRes.Id, err)
		return
	}
	defer pool.ReleaseMessage(resp)
	err = client.notifyContentChanged(obsRes, resp)
	if err != nil {
		// cloud is unsynchronized against device. To recover cloud state, client need to reconnect to cloud.
		log.Errorf("DeviceId: %v, ResourceId: %v: cannot get resource content: %v", obsRes.DeviceId, obsRes.Id, err)
		client.Close()
	}
	if resp.Code() == coapCodes.NotFound {
		client.unpublishResources(ctx, []*pbRA.Resource{obsRes})
	}
}

func (client *Client) addObservedResourceLocked(ctx context.Context, res *pbRA.Resource) error {
	var observation *tcp.Observation
	obs := isObservable(res)
	log.Debugf("DeviceId: %v, ResourceId: %v: Observable: %v: Client.addObservedResourceLocked", res.DeviceId, res.Href, obs)

	if res.Id == cqrsRA.MakeResourceId(res.DeviceId, cloud.StatusHref) {
		return nil
	}

	obsRes := res.Clone()
	if obs {
		obs, err := client.coapConn.Observe(ctx, res.Href, func(req *pool.Message) {
			defer pool.ReleaseMessage(req)
			err := client.notifyContentChanged(obsRes, req)
			if err != nil {
				// cloud is unsynchronized against device. To recover cloud state, client need to reconnect to cloud.
				log.Errorf("DeviceId: %v, ResourceId: %v: cannot get resource content: %v", obsRes.DeviceId, obsRes.Id, err)
				client.Close()
			}
			if req.Code() == coapCodes.NotFound {
				client.unpublishResources(req.Context(), []*pbRA.Resource{obsRes})
			}
		})
		if err != nil {
			log.Errorf("DeviceId: %v, ResourceId: %v: cannot observe resource: %v", obsRes.DeviceId, obsRes.Id, err)
		} else {
			observation = obs
		}
	} else {
		go client.getResourceContent(ctx, obsRes)
	}
	client.observedResources[res.DeviceId][res.InstanceId] = observedResource{res: obsRes, observation: observation}
	return nil
}

func (client *Client) getObservedResources(deviceID string, instanceIDs []int64, matches []*pbRA.Resource) []*pbRA.Resource {
	client.observedResourcesLock.Lock()
	defer client.observedResourcesLock.Unlock()

	getAllDeviceIDMatches := len(instanceIDs) == 0

	if deviceResourcesMap, ok := client.observedResources[deviceID]; ok {
		if getAllDeviceIDMatches {
			for _, value := range deviceResourcesMap {
				matches = append(matches, value.res)
			}
		} else {
			for _, instanceID := range instanceIDs {
				if resource, ok := deviceResourcesMap[instanceID]; ok {
					matches = append(matches, resource.res)
				}
			}
		}
	}

	return matches
}

func (client *Client) removeResource(deviceID string, instanceID int64) {
	if device, ok := client.observedResources[deviceID]; ok {
		delete(device, instanceID)
		if len(client.observedResources[deviceID]) == 0 {
			delete(client.observedResources, deviceID)
		}
	}
}

func (client *Client) popObservation(deviceID string, instanceID int64) *tcp.Observation {
	log.Debugf("remove published resource ocf://%v/%v", deviceID, instanceID)

	var obs *tcp.Observation
	if device, ok := client.observedResources[deviceID]; ok {
		if res, ok := device[instanceID]; ok {
			obs = res.observation
			res.observation = nil
		}
	}

	return obs
}

func (client *Client) unobserveResources(ctx context.Context, rscs []*pbRA.Resource, rscsUnpublished map[string]bool) {
	observartions := client.unobserveAndRemoveResources(rscs, rscsUnpublished)
	for _, obs := range observartions {
		obs.Cancel(ctx)
	}
}

// Close closes coap connection
func (client *Client) Close() error {
	err := client.coapConn.Close()
	if err != nil {
		return fmt.Errorf("cannot close client: %v", err)
	}
	return nil
}

func (client *Client) unobserveAndRemoveResources(rscs []*pbRA.Resource, rscsUnpublished map[string]bool) []*tcp.Observation {
	observartions := make([]*tcp.Observation, 0, 32)

	client.observedResourcesLock.Lock()
	defer client.observedResourcesLock.Unlock()

	for _, resource := range rscs {
		if del, ok := rscsUnpublished[resource.Id]; ok && del {
			log.Debugf("DeviceId: %v, ResourceId: %v: delete resource", resource.DeviceId, resource.Id)
		} else {
			log.Debugf("DeviceId: %v, ResourceId: %v: unobserve resource", resource.DeviceId, resource.Id)
		}

		obs := client.popObservation(resource.DeviceId, resource.InstanceId)
		if obs != nil {
			observartions = append(observartions, obs)
		}
		if rscsUnpublished[resource.Id] {
			client.removeResource(resource.DeviceId, resource.InstanceId)
		}
	}
	return observartions
}

// cleanObservedResourcesOfDevices remove all pbRA observations requested from device to cloud.
func (client *Client) cleanObservedResourcesOfDevices() {
	obs, _ := client.server.observeResourceContainer.PopByRemoteAddr(client.remoteAddrString())
	for _, ob := range obs {
		err := client.server.projection.Unregister(ob.deviceID)
		if err != nil {
			log.Errorf("cannot unregister observed device %v from projection: %v", ob.deviceID, err)
		}
	}
}

func (client *Client) popObservedResources() []*tcp.Observation {
	observartions := make([]*tcp.Observation, 0, 32)
	client.observedResourcesLock.Lock()
	defer client.observedResourcesLock.Unlock()
	for deviceID, instanceIDs := range client.observedResources {
		for instanceID := range instanceIDs {
			obs := client.popObservation(deviceID, instanceID)
			if obs != nil {
				observartions = append(observartions, obs)
			}
			client.removeResource(deviceID, instanceID)
		}
		client.server.clientContainerByDeviceID.Remove(deviceID)
		err := client.server.projection.Unregister(deviceID)
		if err != nil {
			log.Errorf("DeviceId %v: cannot unregister device from projection: %v", deviceID, err)
		}
	}
	return observartions
}

// cleanObservedResources remove all device pbRA observation requested by cloud.
func (client *Client) cleanObservedResources() {
	for _, obs := range client.popObservedResources() {
		obs.Cancel(client.coapConn.Context())
	}
}

// OnClose action when coap connection was closed.
func (client *Client) OnClose() {
	log.Debugf("close client %v", client.coapConn.RemoteAddr())

	atomic.StoreInt32(&client.isClosed, 1)
	client.server.oicPingCache.Delete(client.remoteAddrString())

	client.cleanObservedResources()

	client.cleanObservedResourcesOfDevices()

	authCtx := client.loadAuthorizationContext()

	if client.authCtx.DeviceId != "" {
		err := client.UpdateCloudDeviceStatus(kitNetGrpc.CtxWithToken(context.Background(), authCtx.AccessToken), authCtx.DeviceId, authCtx.AuthorizationContext, false)
		if err != nil {
			// Device will be still reported as online and it can fix his state by next calls online, offline commands.
			log.Errorf("DeviceId %v: cannot handle sign out: cannot update cloud device status: %v", authCtx.DeviceId, err)
		}
	}
}

func (client *Client) storeAuthorizationContext(authCtx authCtx) (oldDeviceId string) {
	log.Debugf("Authorization context stored for client %v, device %v, user %v", client.coapConn.RemoteAddr(), authCtx.GetDeviceId(), authCtx.GetUserId())
	client.authContextLock.Lock()
	defer client.authContextLock.Unlock()
	oldAuthContext := client.authCtx
	client.authCtx = authCtx
	return oldAuthContext.DeviceId
}

func (client *Client) loadAuthorizationContext() authCtx {
	client.authContextLock.Lock()
	defer client.authContextLock.Unlock()
	return client.authCtx
}

func (client *Client) notifyContentChanged(res *pbRA.Resource, notification *pool.Message) error {
	decodeMsgToDebug(client, notification, "RECEIVED-NOTIFICATION")
	authCtx := client.loadAuthorizationContext()

	request := coapconv.MakeNotifyResourceChangedRequest(res.Id, authCtx.AuthorizationContext, client.remoteAddrString(), notification)
	_, err := client.server.raClient.NotifyResourceChanged(kitNetGrpc.CtxWithToken(notification.Context(), authCtx.AccessToken), &request)
	if err != nil {
		return fmt.Errorf("cannot notify resource content changed: %v", err)
	}
	return nil
}

func (client *Client) updateContent(ctx context.Context, resource *pbRA.Resource, reqContentUpdate *raEvents.ResourceUpdatePending) error {
	if reqContentUpdate.AuditContext == nil {
		return fmt.Errorf("invalid AuditContext of resource update pending request")
	}
	if resource.Id == cqrsRA.MakeResourceId(resource.DeviceId, cloud.StatusHref) {
		authCtx := client.loadAuthorizationContext()
		msg := pool.AcquireMessage(ctx)
		msg.SetCode(coapCodes.MethodNotAllowed)
		msg.SetSequence(client.coapConn.Sequence())
		defer pool.ReleaseMessage(msg)

		request := coapconv.MakeConfirmResourceUpdateRequest(resource.Id, reqContentUpdate.AuditContext.CorrelationId, authCtx.AuthorizationContext, client.remoteAddrString(), msg)
		_, err := client.server.raClient.ConfirmResourceUpdate(kitNetGrpc.CtxWithToken(ctx, authCtx.AccessToken), &request)
		if err != nil {
			return err
		}
		return nil
	}

	ctx, cancel := context.WithTimeout(ctx, client.server.RequestTimeout)
	defer cancel()
	req, err := coapconv.NewCoapResourceUpdateRequest(ctx, resource.Href, &reqContentUpdate.ResourceUpdatePending)
	if err != nil {
		return err
	}
	defer pool.ReleaseMessage(req)

	resp, err := client.coapConn.Do(req)
	if err != nil {
		return err
	}
	defer pool.ReleaseMessage(resp)

	decodeMsgToDebug(client, resp, "RESOURCE-UPDATE")

	if resp.Code() == coapCodes.NotFound {
		client.unpublishResources(ctx, []*pbRA.Resource{resource})
	}

	authCtx := client.loadAuthorizationContext()

	request := coapconv.MakeConfirmResourceUpdateRequest(resource.Id, reqContentUpdate.AuditContext.CorrelationId, authCtx.AuthorizationContext, client.remoteAddrString(), resp)
	_, err = client.server.raClient.ConfirmResourceUpdate(kitNetGrpc.CtxWithToken(ctx, authCtx.AccessToken), &request)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) retrieveContent(ctx context.Context, resource *pbRA.Resource, reqContentUpdate *raEvents.ResourceRetrievePending) error {
	if reqContentUpdate.AuditContext == nil {
		return fmt.Errorf("invalid AuditContext of resource retrieve pending request")
	}
	if resource.Id == cqrsRA.MakeResourceId(resource.DeviceId, cloud.StatusHref) {
		authCtx := client.loadAuthorizationContext()
		msg := pool.AcquireMessage(ctx)
		msg.SetCode(coapCodes.Content)
		msg.SetSequence(client.coapConn.Sequence())
		defer pool.ReleaseMessage(msg)

		request := coapconv.MakeConfirmResourceRetrieveRequest(resource.Id, reqContentUpdate.AuditContext.CorrelationId, authCtx.AuthorizationContext, client.remoteAddrString(), msg)
		_, err := client.server.raClient.ConfirmResourceRetrieve(kitNetGrpc.CtxWithToken(ctx, authCtx.AccessToken), &request)
		if err != nil {
			return err
		}
		return nil
	}

	ctx, cancel := context.WithTimeout(ctx, client.server.RequestTimeout)
	defer cancel()
	req, err := coapconv.NewCoapResourceRetrieveRequest(ctx, resource.Href, &reqContentUpdate.ResourceRetrievePending)
	if err != nil {
		return err
	}
	defer pool.ReleaseMessage(req)

	resp, err := client.coapConn.Do(req)
	if err != nil {
		return err
	}
	defer pool.ReleaseMessage(resp)

	decodeMsgToDebug(client, resp, "RESOURCE-RETRIEVE")

	if resp.Code() == coapCodes.NotFound {
		client.unpublishResources(ctx, []*pbRA.Resource{resource})
	}

	authCtx := client.loadAuthorizationContext()
	request := coapconv.MakeConfirmResourceRetrieveRequest(resource.Id, reqContentUpdate.AuditContext.CorrelationId, authCtx.AuthorizationContext, client.remoteAddrString(), resp)
	_, err = client.server.raClient.ConfirmResourceRetrieve(kitNetGrpc.CtxWithToken(ctx, authCtx.AccessToken), &request)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) publishResource(ctx context.Context, resource *pbRA.Resource, ttl int32, connectionId string, sequence uint64, authCtx pbCQRS.AuthorizationContext) (*pbRA.Resource, error) {
	if resource.DeviceId == "" {
		return resource, fmt.Errorf("cannot send command publish resource: invalid DeviceId")
	}
	resource.Href = fixHref(resource.Href)

	if resource.Href == "" || resource.Href == "/" {
		return resource, fmt.Errorf("cannot send command publish resource: invalid Href")
	}
	resource.Id = resource2UUID(resource.DeviceId, resource.Href)

	request := pbRA.PublishResourceRequest{
		AuthorizationContext: &authCtx,
		ResourceId:           resource.Id,
		Resource:             resource,
		TimeToLive:           ttl,
		CommandMetadata: &pbCQRS.CommandMetadata{
			Sequence:     sequence,
			ConnectionId: connectionId,
		},
	}

	response, err := client.server.raClient.PublishResource(ctx, &request)
	if err != nil {
		return resource, fmt.Errorf("cannot process command publish resource: %v", err)
	}

	resource.InstanceId = response.InstanceId
	return resource, nil
}

func (client *Client) unpublishResource(ctx context.Context, resource *pbRA.Resource, authCtx pbCQRS.AuthorizationContext, rscsUnpublished map[string]bool) map[string]bool {
	_, err := client.server.raClient.UnpublishResource(ctx, &pbRA.UnpublishResourceRequest{
		AuthorizationContext: &authCtx,
		ResourceId:           resource.Id,
		CommandMetadata: &pbCQRS.CommandMetadata{
			ConnectionId: client.remoteAddrString(),
			Sequence:     client.coapConn.Sequence(),
		},
	})
	if err != nil {
		// unpublish resource is not critical -> resource is still accessible,
		// but next update of device resource will returns 'not found; and it triggers again unpublish.
		log.Errorf("DeviceId: %v, ResourceId: %v: cannot unpublish resource: %v", resource.DeviceId, resource.Id, err)
		rscsUnpublished[resource.Id] = false
		return rscsUnpublished
	}

	rscsUnpublished[resource.Id] = true
	return rscsUnpublished
}

func (client *Client) unpublishResources(ctx context.Context, rscs []*pbRA.Resource) {
	rscsUnpublished := make(map[string]bool, 32)
	authCtx := client.loadAuthorizationContext()

	for _, resource := range rscs {
		rscsUnpublished = client.unpublishResource(kitNetGrpc.CtxWithToken(ctx, authCtx.AccessToken), resource, authCtx.AuthorizationContext, rscsUnpublished)
	}

	client.unobserveResources(ctx, rscs, rscsUnpublished)
}
