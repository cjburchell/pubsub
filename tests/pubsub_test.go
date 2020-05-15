package tests

import (
	"context"
	"testing"

	"github.com/cjburchell/pubsub"
	"github.com/stretchr/testify/assert"
)

func TestPublishMemory(t *testing.T) {

	ctx := context.Background()

	p, err := pubsub.Create(ctx, pubsub.Settings{Provider: pubsub.MemoryProvider})
	assert.Nil(t, err)
	c := make(chan []byte)
	sub, err := p.SubscribeChan(ctx, "test", c)
	assert.Nil(t, err)
	assert.NotNil(t, sub)

	err = p.Publish(ctx, "test", []byte("this is a test"))
	assert.Nil(t, err)

	err = sub.Close()
	assert.Nil(t, err)
	result := <-c
	assert.Equal(t, []byte("this is a test"), result)
}

func TestBadProvider(t *testing.T) {
	ctx := context.Background()
	p, err := pubsub.Create(ctx, pubsub.Settings{Provider: "test"})
	assert.NotNil(t, err)
	assert.Nil(t, p)
}
