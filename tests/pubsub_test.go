package tests

import (
	"github.com/cjburchell/pubsub"
	"github.com/cjburchell/pubsub/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestPublishMemory(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mocks.NewMockISettings(ctrl)
	m.EXPECT().Get(gomock.Eq("PubSubProvider"), gomock.Any()).Return("memory").AnyTimes()

	p, err := pubsub.Create(m)
	assert.Nil(t, err)
	c := make(chan []byte)
	sub, err := p.SubscribeChan("test", c)
	assert.Nil(t, err)
	assert.NotNil(t, sub)

	err = p.Publish("test", []byte("this is a test"))
	assert.Nil(t, err)

	err = sub.Close()
	assert.Nil(t, err)
    result := <- c
	assert.Equal(t, []byte("this is a test"), result)
}

func TestBadProvider(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mocks.NewMockISettings(ctrl)
	m.EXPECT().Get(gomock.Eq("PubSubProvider"), gomock.Any()).Return("test").AnyTimes()

	p, err := pubsub.Create(m)
	assert.NotNil(t, err)
	assert.Nil(t, p)
}