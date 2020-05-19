package pubsub

import (
	"context"

	"github.com/nats-io/nats.go"
)

// createNatsProvider create the provider
func createNatsProvider(settings NatsProviderSettings) (provider, error) {
	var loginOption nats.Option
	if settings.Token != ""{
		loginOption = nats.Token(settings.Token)
	} else if settings.User != "" &&  settings.Password != "" {
		loginOption = nats.UserInfo(settings.User, settings.Password)
	}

	natsConn, err := nats.Connect(
		settings.URL,
		loginOption)

	if err != nil {
		return nil, err
	}

	return nProvider{natsConn: natsConn}, nil
}

type nProvider struct {
	natsConn *nats.Conn
}

type natsSubscription struct {
	natsSub *nats.Subscription
}

func (n natsSubscription) Close() error {
	return n.natsSub.Unsubscribe()
}

func (n nProvider) Publish(_ context.Context, topic string, msg []byte) error {
	return n.natsConn.Publish(topic, msg)
}

func (n nProvider) Subscribe(_ context.Context, topic string, handler MsgHandler) (ISubscription, error) {
	sub, err := n.natsConn.Subscribe(topic, func(msg *nats.Msg) {
		handler(msg.Data)
	})

	if err != nil {
		return nil, err
	}

	return natsSubscription{natsSub: sub}, nil
}
