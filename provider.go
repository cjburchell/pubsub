package pubsub

import (
	"context"
	"fmt"
)

type provider interface {
	Publish(ctx context.Context, topic string, msg []byte) error
	Subscribe(ctx context.Context, topic string, handler MsgHandler) (ISubscription, error)
}

type providerType string

const (
	googlePubSubProvider providerType = "googlePubSub"
	natsProvider                 providerType = "nats"
	memoryProvider               providerType = "memory"
)

func getProvider(ctx context.Context, providerType providerType, settings ISettings) (provider, error){
	switch providerType {
	case googlePubSubProvider:
		return createGoogleProvider(ctx, settings)
	case natsProvider:
		return createNatsProvider(settings)
	case memoryProvider:
		return createMemoryProvider()
	}

	return nil, fmt.Errorf("unknown povider %s", providerType)
}
