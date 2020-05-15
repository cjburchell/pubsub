package pubsub

import "github.com/cjburchell/settings-go"

// Settings for the pub sub
type Settings struct {
	Provider ProviderType
	Google   GoogleProviderSettings
	Nats     NatsProviderSettings
}

// NatsProviderSettings type
type NatsProviderSettings struct {
	URL      string
	Token    string
	User     string
	Password string
}

// GoogleProviderSettings type
type GoogleProviderSettings struct {
	ProjectID          string
	CredentialsFile    string
	SubscriptionSuffix string
	CreateTopic        bool
	CreateSubscription bool
}

// GetSettings Get the pub sub settings
func GetSettings(settings settings.ISettings) Settings {
	return Settings{
		Provider: ProviderType(settings.Get("PubSubProvider", string(MemoryProvider))),
		Google:   getGoogleSettings(settings),
		Nats:     getNatsSettings(settings),
	}
}

func getNatsSettings(settings settings.ISettings) NatsProviderSettings {
	return NatsProviderSettings{
		URL:      settings.Get("pubSubNatsUrl", "tcp://nats:4222"),
		Token:    settings.Get("pubSubNatsToken", "token"),
		User:     settings.Get("pubSubNatsUser", "admin"),
		Password: settings.Get("pubSubNatsPassword", "password"),
	}
}

func getGoogleSettings(settings settings.ISettings) GoogleProviderSettings {
	return GoogleProviderSettings{
		ProjectID:          settings.Get("googlePubSubProjectId", ""),
		CredentialsFile:    settings.Get("googlePubSubCredentialsFile", ""),
		SubscriptionSuffix: settings.Get("googlePubSubscriptionSuffix", ""),
		CreateTopic:        settings.GetBool("googlePubSubCreateTopic", false),
		CreateSubscription: settings.GetBool("googlePubSubCreateSubscription", false),
	}
}
