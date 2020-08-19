package events

import (
	"github.com/plgd-dev/cloud/resource-aggregate/pb"
	"github.com/plgd-dev/kit/net/http"
)

type ResourceChanged struct {
	pb.ResourceChanged
}

func (e ResourceChanged) Version() uint64 {
	return e.EventMetadata.Version
}

func (e ResourceChanged) Marshal() ([]byte, error) {
	return e.ResourceChanged.Marshal()
}

func (e *ResourceChanged) Unmarshal(b []byte) error {
	return e.ResourceChanged.Unmarshal(b)
}

func (e ResourceChanged) EventType() string {
	return http.ProtobufContentType(&pb.ResourceChanged{})
}

func (e ResourceChanged) AggregateId() string {
	return e.Id
}
