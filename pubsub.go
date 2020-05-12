package pubsub

// IPubSub Interface
type IPubSub interface {
	Publish(topic string, msg []byte) error
	Subscribe(topic string, handler MsgHandler) (ISubscription, error)
	SubscribeChan(topic string, channel chan []byte) (ISubscription, error)
}

// ISubscription interface
type ISubscription interface {
	Close() error
}

// Create the Pub sub
func Create(settings ISettings) (IPubSub, error) {
	ptype := settings.Get("PubSubProvider", string(memoryProvider))
	p, err := getProvider(providerType(ptype), settings)
	if err != nil{
		return nil, err
	}
	return &pubSub{p}, nil
}

type pubSub struct {
	provider provider
}

func (p* pubSub) Publish(topic string, msg []byte) error {
	return p.provider.Publish(topic, msg)
}

func (p* pubSub) Subscribe(topic string, handler MsgHandler) (ISubscription, error) {
	return p.provider.Subscribe(topic, handler)
}

func (p* pubSub) SubscribeChan(	topic string, channel chan []byte) (ISubscription, error) {
	return p.provider.Subscribe(topic, func(msg []byte) {
		channel <- msg
	})
}