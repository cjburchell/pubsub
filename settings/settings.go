package settings

import (
	"github.com/cjburchell/pubsub"
	"github.com/cjburchell/settings-go"
)

// Get the pub sub settings
func Get(settings settings.ISettings) pubsub.Settings {
	return pubsub.Settings{
		Provider: pubsub.ProviderType(settings.Get("PubSubProvider", string(pubsub.MemoryProvider))),
		Google:   getGoogleSettings(settings),
		Nats:     getNatsSettings(settings),
	}
}

func getNatsSettings(settings settings.ISettings) pubsub.NatsProviderSettings {
	return pubsub.NatsProviderSettings{
		URL:      settings.Get("pubSubNatsUrl", "tcp://nats:4222"),
		Token:    settings.Get("pubSubNatsToken", "token"),
		User:     settings.Get("pubSubNatsUser", "admin"),
		Password: settings.Get("pubSubNatsPassword", "password"),
	}
}

func getGoogleSettings(settings settings.ISettings) pubsub.GoogleProviderSettings {
	return pubsub.GoogleProviderSettings{
		ProjectID:          settings.Get("googlePubSubProjectId", ""),
		CredentialsFile:    settings.Get("googlePubSubCredentialsFile", ""),
		SubscriptionSuffix: settings.Get("googlePubSubscriptionSuffix", ""),
		CreateTopic:        settings.GetBool("googlePubSubCreateTopic", false),
		CreateSubscription: settings.GetBool("googlePubSubCreateSubscription", false),
	}
}
