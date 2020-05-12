package pubsub

import "context"

// createMemoryProvider create the provider
func createMemoryProvider() (provider, error){
	return &memProvider{make(map[string]*session, 0)}, nil
}

type memProvider struct {
	sessions map[string]*session
}

func (m*memProvider) Publish(_ context.Context, topic string, msg []byte) error {
	if session, ok := m.sessions[topic]; ok {
		for _, sub := range session.subscriptions {
			go sub.handler(msg)
		}
	}

	return nil
}

func (m*memProvider) Subscribe(_ context.Context, topic string, handler MsgHandler) (ISubscription, error) {
	var sub = &subscription{
		subject: topic,
		handler: handler,
		provider: m,
	}
	if s, ok := m.sessions[topic]; ok {
		s.subscriptions = append(s.subscriptions, sub)
	} else {
		var s session
		s.subscriptions = append(s.subscriptions, sub)
		m.sessions[topic] = &s
	}

	return sub, nil
}

type subscription struct {
	subject  string
	handler  MsgHandler
	provider *memProvider
}

func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// Close the subscription
func (sub *subscription) Close() error {
	session := sub.provider.sessions[sub.subject]
	session.subscriptions = removeSubscription(session.subscriptions, sliceIndex(len(session.subscriptions), func(i int) bool { return session.subscriptions[i] == sub }))
	if len(session.subscriptions) == 0 {
		delete(sub.provider.sessions, sub.subject)
	}

	return nil
}

type session struct {
	subscriptions []*subscription
}

func removeSubscription(s []*subscription, i int) []*subscription {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
