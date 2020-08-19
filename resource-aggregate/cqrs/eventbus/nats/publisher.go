package nats

import (
	cqrsUtils "github.com/plgd-dev/cloud/resource-aggregate/cqrs"
	cqrsNats "github.com/plgd-dev/cqrs/eventbus/nats"
)

type Publisher struct {
	*cqrsNats.Publisher
}

// NewPublisher creates new publisher with proto marshaller.
func NewPublisher(config Config, opts ...Option) (*Publisher, error) {
	for _, o := range opts {
		config = o(config)
	}

	p, err := cqrsNats.NewPublisher(config.URL, cqrsUtils.Marshal, config.Options...)
	if err != nil {
		return nil, err
	}
	return &Publisher{
		p,
	}, nil
}

// Close closes the publisher.
func (p *Publisher) Close() {
	p.Publisher.Close()
}
