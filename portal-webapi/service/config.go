package service

import (
	"encoding/json"
	"fmt"

	"github.com/plgd-dev/cloud/pkg/net/grpc"
)

type httpProto string

func (a *httpProto) Decode(value string) error {
	switch value {
	case "http", "https":
		*a = httpProto(value)
		return nil
	default:
		return fmt.Errorf("unsupported protocol type %v", value)
	}
}

//Config represent application configuration
type Config struct {
	grpc.Config
	ResourceDirectoryAddr string `envconfig:"RESOURCE_DIRECTORY_ADDRESS"  default:"127.0.0.1:9100"`
	ResourceAggregateAddr string `envconfig:"RESOURCE_AGGREGATE_ADDRESS"  default:"127.0.0.1:9100"`
}

//String return string representation of Config
func (c Config) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("config: \n%v\n", string(b))
}
