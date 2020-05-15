package pubsub

import (
	"context"
	"fmt"
)

type provider interface {
	Publish(ctx context.Context, topic string, msg []byte) error
	Subscribe(ctx context.Context, topic string, handler MsgHandler) (ISubscription, error)
}

//ProviderType the type of provider
type ProviderType string

const (
	// GooglePubSubProvider type
	GooglePubSubProvider ProviderType = "googlePubSub"
	// NatsProvider type
	NatsProvider ProviderType = "nats"
	// MemoryProvider type
	MemoryProvider ProviderType = "memory"
)

func getProvider(ctx context.Context, settings Settings) (provider, error) {
	switch settings.Provider {

	case GooglePubSubProvider:
		return createGoogleProvider(ctx, settings.Google)
	case NatsProvider:
		return createNatsProvider(settings.Nats)
	case MemoryProvider:
		return createMemoryProvider()
	}

	return nil, fmt.Errorf("unknown povider %s", settings.Provider)
}
