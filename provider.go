package pubsub

import (
	"fmt"
)

type provider interface {
	Publish(topic string, msg []byte) error
	Subscribe(topic string, handler MsgHandler) (ISubscription, error)
}

type providerType string

const (
	googlePubSubProvider providerType = "googlePubSub"
	natsProvider                 providerType = "nats"
	memoryProvider               providerType = "memory"
)

func getProvider(providerType providerType, settings ISettings) (provider, error){
	switch providerType {
	case googlePubSubProvider:
		return createGoogleProvider(settings)
	case natsProvider:
		return createNatsProvider(settings)
	case memoryProvider:
		return createMemoryProvider()
	}

	return nil, fmt.Errorf("unknown povider %s", providerType)
}
