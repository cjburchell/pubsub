package pubsub

import (
	"context"
	"github.com/nats-io/nats.go"
)

// createNatsProvider create the provider
func createNatsProvider(settings ISettings) (provider, error) {
	url := settings.Get("pubSubNatsUrl", "tcp://nats:4222")
	token := settings.Get("pubSubNatsToken", "token")
	user := settings.Get("pubSubNatsUser", "admin")
	password := settings.Get("pubSubNatsPassword", "password")

	natsConn, err := nats.Connect(
		url,
		nats.Token(token),
		nats.UserInfo(user, password))
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

func (n nProvider) Publish(_ context.Context,topic string, msg []byte) error {
	return n.natsConn.Publish(topic, msg)
}

func (n nProvider) Subscribe(_ context.Context, topic string, handler MsgHandler) (ISubscription, error) {
	sub, err := n.natsConn.Subscribe(topic, func(msg *nats.Msg) {
		handler(msg.Data)
	})

	if err != nil{
		return nil, err
	}

	return natsSubscription{natsSub: sub}, nil
}

