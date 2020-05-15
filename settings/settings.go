package settings

import (
	"github.com/cjburchell/pubsub"
	"github.com/cjburchell/settings-go"
)

// Get the pub sub settings
func Get(settings settings.ISettings) pubsub.Settings {
	return pubsub.Settings{
		Provider: pubsub.ProviderType(settings.Get("Provider", string(pubsub.MemoryProvider))),
		Google:   getGoogleSettings(settings.GetSection("Google")),
		Nats:     getNatsSettings(settings.GetSection("Nats")),
	}
}

func getNatsSettings(settings settings.ISettings) pubsub.NatsProviderSettings {
	return pubsub.NatsProviderSettings{
		URL:      settings.Get("Url", "tcp://nats:4222"),
		Token:    settings.Get("Token", "token"),
		User:     settings.Get("User", "admin"),
		Password: settings.Get("Password", "password"),
	}
}

func getGoogleSettings(settings settings.ISettings) pubsub.GoogleProviderSettings {
	return pubsub.GoogleProviderSettings{
		ProjectID:          settings.Get("ProjectId", ""),
		CredentialsFile:    settings.Get("CredentialsFile", ""),
		SubscriptionSuffix: settings.Get("SubscriptionSuffix", ""),
		CreateTopic:        settings.GetBool("CreateTopic", false),
		CreateSubscription: settings.GetBool("CreateSubscription", false),
	}
}
