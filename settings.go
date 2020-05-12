package pubsub

// ISettings interface
type ISettings interface {
	Get(key string, fallback string) string
	GetBool(key string, fallback bool) bool
}

// MsgHandler is a callback function that processes messages delivered to
// asynchronous subscribers.
type MsgHandler func(msg []byte)