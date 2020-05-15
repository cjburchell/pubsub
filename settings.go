package pubsub

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
