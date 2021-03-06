package service

import (
	"context"

	"github.com/plgd-dev/cloud/cloud2cloud-connector/events"
	raEvents "github.com/plgd-dev/cloud/resource-aggregate/events"
	"github.com/plgd-dev/kit/log"
	"github.com/plgd-dev/sdk/schema"
)

type deviceSubsciptionHandler struct {
	subData   *SubscriptionData
	emitEvent emitEventFunc
}

func fixResourceLink(r schema.ResourceLink) schema.ResourceLink {
	r.Href = getHref(r.DeviceID, r.Href)
	r.ID = ""
	return r
}

func (h *deviceSubsciptionHandler) HandleResourcePublished(ctx context.Context, val *raEvents.ResourceLinksPublished) error {
	toSend := make([]schema.ResourceLink, 0, 32)
	for _, l := range val.GetResources() {
		toSend = append(toSend, fixResourceLink(l.ToSchema()))
	}
	if len(toSend) == 0 && len(val.GetResources()) > 0 {
		return nil
	}
	remove, err := h.emitEvent(ctx, events.EventType_ResourcesPublished, h.subData.Data(), h.subData.IncrementSequenceNumber, toSend)
	if err != nil {
		log.Errorf("deviceSubsciptionHandler.HandleResourcePublished: cannot emit event: %v", err)
	}
	if remove {
		return err
	}
	return nil
}

func (h *deviceSubsciptionHandler) HandleResourceUnpublished(ctx context.Context, val *raEvents.ResourceLinksUnpublished) error {
	toSend := make([]schema.ResourceLink, 0, 32)
	for _, l := range val.GetHrefs() {
		toSend = append(toSend, fixResourceLink(schema.ResourceLink{
			DeviceID: val.GetDeviceId(),
			Href:     l,
		}))
	}
	if len(toSend) == 0 && len(val.GetHrefs()) > 0 {
		return nil
	}
	remove, err := h.emitEvent(ctx, events.EventType_ResourcesUnpublished, h.subData.Data(), h.subData.IncrementSequenceNumber, toSend)
	if err != nil {
		log.Errorf("deviceSubsciptionHandler.HandleResourceUnpublished: cannot emit event: %v", err)
	}
	if remove {
		return err
	}
	return nil
}

type resourcePublishedHandler struct {
	h *deviceSubsciptionHandler
}

func (h *resourcePublishedHandler) HandleResourcePublished(ctx context.Context, val *raEvents.ResourceLinksPublished) error {
	return h.h.HandleResourcePublished(ctx, val)
}

type resourceUnpublishedHandler struct {
	h *deviceSubsciptionHandler
}

func (h *resourceUnpublishedHandler) HandleResourceUnpublished(ctx context.Context, val *raEvents.ResourceLinksUnpublished) error {
	return h.h.HandleResourceUnpublished(ctx, val)
}
