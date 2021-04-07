package service

import (
	"fmt"
	"time"

	"github.com/plgd-dev/cloud/pkg/config"
	"github.com/plgd-dev/cloud/pkg/log"
	"github.com/plgd-dev/cloud/pkg/net/grpc/client"
	grpcServer "github.com/plgd-dev/cloud/pkg/net/grpc/server"
	client2 "github.com/plgd-dev/cloud/pkg/security/oauth/manager"
	eventbusConfig "github.com/plgd-dev/cloud/resource-aggregate/cqrs/eventbus/config"
	eventstoreConfig "github.com/plgd-dev/cloud/resource-aggregate/cqrs/eventstore/config"
)

//Config represent application configuration
type Config struct {
	Log     log.Config    `yaml:"log" json:"log"`
	APIs    APIsConfig    `yaml:"apis" json:"apis"`
	Clients ClientsConfig `yaml:"clients" json:"clients"`
}

func (c *Config) Validate() error {
	err := c.APIs.Validate()
	if err != nil {
		return fmt.Errorf("apis.%w", err)
	}
	err = c.Clients.Validate()
	if err != nil {
		return fmt.Errorf("clients.%w", err)
	}
	return nil
}

type APIsConfig struct {
	GRPC grpcServer.Config `yaml:"grpc" json:"grpc"`
}

func (c *APIsConfig) Validate() error {
	err := c.GRPC.Validate()
	if err != nil {
		return fmt.Errorf("grpc.%w", err)
	}
	return nil
}

type EventStoreConfig struct {
	Connection                   eventstoreConfig.Config `yaml:",inline" json:",inline"`
	SnapshotThreshold            int                     `yaml:"snapshotThreshold" json:"snapshotThreshold" default:"16"`
	ConcurrencyExceptionMaxRetry int                     `yaml:"occMaxRetry" json:"occMaxRetry" default:"8"`
}

func (c *EventStoreConfig) Validate() error {
	if c.SnapshotThreshold <= 0 {
		return fmt.Errorf("snapshotThreshold('%v')", c.SnapshotThreshold)
	}
	if c.ConcurrencyExceptionMaxRetry <= 0 {
		return fmt.Errorf("occMaxRetry('%v')", c.ConcurrencyExceptionMaxRetry)
	}
	return c.Connection.Validate()
}

type AuthorizationServerConfig struct {
	Connection      client.Config    `yaml:"grpc" json:"grpc"`
	PullFrequency   time.Duration    `yaml:"pullFrequency" json:"pullFrequency" default:"15s"`
	CacheExpiration time.Duration    `yaml:"cacheExpiration" json:"cacheExpiration" default:"1m"`
	OAuth           client2.ConfigV2 `yaml:"oauth" json:"oauth"`
}

func (c *AuthorizationServerConfig) Validate() error {
	if c.PullFrequency <= 0 {
		return fmt.Errorf("pullFrequency('%v')", c.PullFrequency)
	}
	if c.CacheExpiration <= 0 {
		return fmt.Errorf("cacheExpiration('%v')", c.CacheExpiration)
	}
	err := c.OAuth.Validate()
	if err != nil {
		return fmt.Errorf("oauth.%w", err)
	}
	err = c.Connection.Validate()
	if err != nil {
		return fmt.Errorf("grpc.%w", err)
	}
	return err
}

type ClientsConfig struct {
	Eventbus   eventbusConfig.Config     `yaml:"eventBus" json:"eventBus"`
	Eventstore EventStoreConfig          `yaml:"eventStore" json:"eventStore"`
	AuthServer AuthorizationServerConfig `yaml:"authorizationServer" json:"authorizationServer"`
}

func (c *ClientsConfig) Validate() error {
	err := c.AuthServer.Validate()
	if err != nil {
		return fmt.Errorf("authorizationServer.%w", err)
	}
	err = c.Eventbus.Validate()
	if err != nil {
		return fmt.Errorf("eventbus.%w", err)
	}
	err = c.Eventstore.Validate()
	if err != nil {
		return fmt.Errorf("eventstore.%w", err)
	}
	return nil
}

//String return string representation of Config
func (c Config) String() string {
	return config.ToString(c)
}
