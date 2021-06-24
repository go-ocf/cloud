package events

import (
	"time"

	"google.golang.org/protobuf/proto"
)

const eventTypeResourceUpdatePending = "ocf.cloud.resourceaggregate.events.resourceupdatepending"

func (e *ResourceUpdatePending) Version() uint64 {
	return e.GetEventMetadata().GetVersion()
}

func (e *ResourceUpdatePending) Marshal() ([]byte, error) {
	return proto.Marshal(e)
}

func (e *ResourceUpdatePending) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, e)
}

func (e *ResourceUpdatePending) EventType() string {
	return eventTypeResourceUpdatePending
}

func (e *ResourceUpdatePending) AggregateID() string {
	return e.GetResourceId().ToUUID()
}

func (e *ResourceUpdatePending) GroupID() string {
	return e.GetResourceId().GetDeviceId()
}

func (e *ResourceUpdatePending) IsSnapshot() bool {
	return false
}

func (e *ResourceUpdatePending) Timestamp() time.Time {
	return time.Unix(0, e.GetEventMetadata().GetTimestamp())
}
