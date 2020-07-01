package projection

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-ocf/cqrs"
	"github.com/go-ocf/cqrs/eventbus"
	"github.com/go-ocf/cqrs/eventstore"
	"github.com/go-ocf/kit/log"

	raCqrsUtils "github.com/go-ocf/cloud/resource-aggregate/cqrs"
)

// Projection projects events from resource aggregate.
type Projection struct {
	projection *projection
}

// NewProjection creates new resource projection.
func NewProjection(ctx context.Context, name string, store eventstore.EventStore, subscriber eventbus.Subscriber, factoryModel eventstore.FactoryModelFunc) (*Projection, error) {
	projection, err := newProjection(ctx, name, store, subscriber, factoryModel, raCqrsUtils.GetTopics)
	if err != nil {
		return nil, fmt.Errorf("cannot create resource projection: %w", err)
	}
	return &Projection{projection: projection}, nil
}

// Register registers deviceId, loads events from eventstore and subscribe to eventbus.
// It can be called multiple times for same deviceId but after successful the a call Unregister
// must be called same times to free resources.
func (p *Projection) Register(ctx context.Context, deviceId string) (created bool, err error) {
	return p.projection.register(ctx, deviceId, []eventstore.SnapshotQuery{{GroupId: deviceId}})
}

// Unregister unregisters device and his resource from projection.
func (p *Projection) Unregister(deviceId string) error {
	return p.projection.unregister(deviceId)
}

// Models returns models for device, resource or nil for non exist.
func (p *Projection) Models(deviceId, resourceId string) []eventstore.Model {
	return p.projection.models([]eventstore.SnapshotQuery{{GroupId: deviceId, AggregateId: resourceId}})
}

// ForceUpdate invokes update registered resource model from evenstore.
func (p *Projection) ForceUpdate(ctx context.Context, deviceId, resourceId string) error {
	err := p.projection.forceUpdate(ctx, deviceId, []eventstore.SnapshotQuery{{GroupId: deviceId, AggregateId: resourceId}})
	if err != nil {
		return fmt.Errorf("cannot force update resource projection: %w", err)
	}
	return err
}

type projection struct {
	cqrsProjection *cqrs.Projection

	topicManager *TopicManager
	refCountMap  *Map
}

func newProjection(ctx context.Context, name string, store eventstore.EventStore, subscriber eventbus.Subscriber, factoryModel eventstore.FactoryModelFunc, getTopics GetTopicsFunc) (*projection, error) {
	cqrsProjection, err := cqrs.NewProjection(ctx, store, name, subscriber, factoryModel, log.Debugf)
	if err != nil {
		return nil, fmt.Errorf("cannot create Projection: %w", err)
	}
	return &projection{
		cqrsProjection: cqrsProjection,
		topicManager:   NewTopicManager(getTopics),
		refCountMap:    NewMap(),
	}, nil
}

// ForceUpdate invokes update registered resource model from evenstore.
func (p *projection) forceUpdate(ctx context.Context, deviceID string, query []eventstore.SnapshotQuery) error {
	v, ok := p.refCountMap.LoadWithFunc(deviceID, func(v interface{}) interface{} {
		r := v.(*RefCounter)
		r.Acquire()
		return r
	})
	if !ok {
		return fmt.Errorf("cannot force update projection for %v: not found", deviceID)
	}
	r := v.(*RefCounter)
	defer r.Release(ctx)
	d := r.Data().(*deviceProjection)
	d.mutex.Lock()
	defer d.mutex.Unlock()

	err := p.cqrsProjection.Project(ctx, query)
	if err != nil {
		return fmt.Errorf("cannot force update projection for %v: %w", deviceID, err)
	}
	return nil
}

func (p *projection) models(query []eventstore.SnapshotQuery) []eventstore.Model {
	for _, q := range query {
		if q.GroupId != "" {
			v, ok := p.refCountMap.LoadWithFunc(q.GroupId, func(v interface{}) interface{} {
				r := v.(*RefCounter)
				r.Acquire()
				return r
			})
			if !ok {
				continue
			}
			r := v.(*RefCounter)
			defer r.Release(context.Background())
			d := r.Data().(*deviceProjection)
			d.mutex.Lock()
			defer d.mutex.Unlock()
		}
	}
	return p.cqrsProjection.Models(query)
}

type deviceProjection struct {
	mutex sync.Mutex
}

func (p *projection) register(ctx context.Context, deviceID string, query []eventstore.SnapshotQuery) (created bool, err error) {
	v, loaded := p.refCountMap.LoadOrStoreWithFunc(deviceID, func(v interface{}) interface{} {
		r := v.(*RefCounter)
		r.Acquire()
		return r
	}, func() interface{} {
		return NewRefCounter(&deviceProjection{}, func(ctx context.Context, data interface{}) error {
			p.refCountMap.Delete(deviceID)
			topics, updateSubscriber := p.topicManager.Remove(deviceID)
			if updateSubscriber {
				err := p.cqrsProjection.SubscribeTo(topics)
				if err != nil {
					log.Errorf("cannot change topics for projection device %v: %w", deviceID, err)
				}
			}
			return p.cqrsProjection.Forget([]eventstore.SnapshotQuery{
				{GroupId: deviceID},
			})
		})
	})
	r := v.(*RefCounter)
	d := r.Data().(*deviceProjection)
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if loaded {
		return false, nil
	}
	topics, updateSubscriber := p.topicManager.Add(deviceID)
	if updateSubscriber {
		err := p.cqrsProjection.SubscribeTo(topics)
		if err != nil {
			r.Release(ctx)
			return false, fmt.Errorf("cannot register device %v: %w", deviceID, err)
		}
	}

	err = p.cqrsProjection.Project(ctx, query)
	if err != nil {
		r.Release(ctx)
		return false, fmt.Errorf("cannot register device %v: %w", deviceID, err)
	}

	return true, nil
}

func (p *projection) unregister(deviceID string) error {
	v, ok := p.refCountMap.LoadWithFunc(deviceID, func(v interface{}) interface{} {
		r := v.(*RefCounter)
		r.Acquire()
		return r
	})
	if !ok {
		return fmt.Errorf("cannot unregister projection for %v: not found", deviceID)
	}
	r := v.(*RefCounter)
	d := r.Data().(*deviceProjection)
	d.mutex.Lock()
	defer d.mutex.Unlock()
	ctx := context.Background()
	r.Release(ctx)
	return r.Release(ctx)
}
