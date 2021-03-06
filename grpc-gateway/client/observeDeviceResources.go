package client

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/plgd-dev/cloud/resource-aggregate/events"
)

type DeviceResourcesObservationHandler = interface {
	HandleResourcePublished(ctx context.Context, val *events.ResourceLinksPublished) error
	HandleResourceUnpublished(ctx context.Context, val *events.ResourceLinksUnpublished) error
	OnClose()
	Error(err error)
}

type deviceResourcesObservation struct {
	h                  DeviceResourcesObservationHandler
	removeSubscription func()
}

func (o *deviceResourcesObservation) HandleResourcePublished(ctx context.Context, val *events.ResourceLinksPublished) error {
	return o.h.HandleResourcePublished(ctx, val)
}

func (o *deviceResourcesObservation) HandleResourceUnpublished(ctx context.Context, val *events.ResourceLinksUnpublished) error {
	return o.h.HandleResourceUnpublished(ctx, val)
}

func (o *deviceResourcesObservation) OnClose() {
	o.removeSubscription()
	o.h.OnClose()
}
func (o *deviceResourcesObservation) Error(err error) {
	o.removeSubscription()
	o.h.Error(err)
}

func (c *Client) ObserveDeviceResources(ctx context.Context, deviceID string, handler DeviceResourcesObservationHandler) (string, error) {
	ID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	sub, err := c.NewDeviceSubscription(ctx, deviceID, &deviceResourcesObservation{
		h: handler,
		removeSubscription: func() {
			c.stopObservingDeviceResources(ID.String())
		},
	})
	if err != nil {
		return "", err
	}
	c.insertSubscription(ID.String(), sub)

	return ID.String(), err
}

func (c *Client) stopObservingDeviceResources(observationID string) (wait func(), err error) {
	s, err := c.popSubscription(observationID)
	if err != nil {
		return nil, err
	}
	return s.Cancel()
}

func (c *Client) StopObservingDeviceResources(ctx context.Context, observationID string) error {
	wait, err := c.stopObservingDeviceResources(observationID)
	if err != nil {
		return err
	}
	wait()
	return nil
}
