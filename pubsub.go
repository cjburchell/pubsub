package pubsub

import (
	"context"
)

// IPubSub Interface
type IPubSub interface {
	Publish(ctx context.Context, topic string, msg []byte) error
	Subscribe(ctx context.Context, topic string, handler MsgHandler) (ISubscription, error)
	SubscribeChan(ctx context.Context, topic string, channel chan []byte) (ISubscription, error)
}

// ISubscription interface
type ISubscription interface {
	Close() error
}

// MsgHandler is a callback function that processes messages delivered to
// asynchronous subscribers.
type MsgHandler func(msg []byte)

// Create the Pub sub
func Create(ctx context.Context, settings Settings) (IPubSub, error) {
	p, err := getProvider(ctx, settings)
	if err != nil {
		return nil, err
	}
	return &pubSub{p}, nil
}

type pubSub struct {
	provider provider
}

func (p *pubSub) Publish(ctx context.Context, topic string, msg []byte) error {
	return p.provider.Publish(ctx, topic, msg)
}

func (p *pubSub) Subscribe(ctx context.Context, topic string, handler MsgHandler) (ISubscription, error) {
	return p.provider.Subscribe(ctx, topic, handler)
}

func (p *pubSub) SubscribeChan(ctx context.Context, topic string, channel chan []byte) (ISubscription, error) {
	return p.provider.Subscribe(ctx, topic, func(msg []byte) {
		channel <- msg
	})
}
