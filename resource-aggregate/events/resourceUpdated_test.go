package events_test

import (
	"testing"

	commands "github.com/plgd-dev/cloud/resource-aggregate/commands"
	"github.com/plgd-dev/cloud/resource-aggregate/events"
	"github.com/plgd-dev/go-coap/v2/message"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestResourceUpdated_CopyData(t *testing.T) {
	evt := events.ResourceUpdated{
		ResourceId: &commands.ResourceId{
			DeviceId: "dev1",
			Href:     "/dev1",
		},
		Status: commands.Status_ACCEPTED,
		Content: &commands.Content{
			Data:              []byte{'t', 'e', 'x', 't'},
			ContentType:       "text",
			CoapContentFormat: int32(message.TextPlain),
		},
		AuditContext: &commands.AuditContext{
			UserId:        "501",
			CorrelationId: "1",
		},
		EventMetadata: &events.EventMetadata{
			Version:      42,
			Timestamp:    12345,
			ConnectionId: "con1",
			Sequence:     1,
		},
	}

	type args struct {
		event *events.ResourceUpdated
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Identity",
			args: args{
				event: &evt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e events.ResourceUpdated
			e.CopyData(tt.args.event)
			require.True(t, proto.Equal(tt.args.event, &e))
		})
	}
}
