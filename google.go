package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"google.golang.org/api/option"
)

func createGoogleProvider(ctx context.Context, settings ISettings) (provider, error) {
	projectID := settings.Get("googlePubSubProjectId", "")
	if projectID == "" {
		return nil, fmt.Errorf("missing setting googlePubSubProjectId")
	}

	credentialsFile := settings.Get("googlePubSubCredentialsFile", "")
	if credentialsFile == "" {
		return nil, fmt.Errorf("missing setting googlePubSubCredentialsFile")
	}

	subscriptionSuffix := settings.Get("googlePubSubscriptionSuffix", "")
	if credentialsFile == "" {
		return nil, fmt.Errorf("missing setting googlePubSubscriptionSuffix")
	}

	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, err
	}

	return &googleProvider{
		client:             client,
		projectID:          projectID,
		subscriptionSuffix: subscriptionSuffix,
		createTopic:        settings.GetBool("googlePubSubCreateTopic", false),
		createSubscription: settings.GetBool("googlePubSubCreateSubscription", false),
	}, nil
}

type googleProvider struct {
	createTopic        bool
	createSubscription bool
	subscriptionSuffix string
	projectID          string
	client             *pubsub.Client
}

func (g*googleProvider) Publish(ctx context.Context,topicID string, msg []byte) error {
	topic := g.client.Topic(topicID)
	if ok, err := topic.Exists(ctx); !ok || err != nil {
		if err != nil{
			return err
		}
		return fmt.Errorf("missing topic %s", topicID)
	}

	res := topic.Publish(ctx, &pubsub.Message{Data: msg})
	_, err := res.Get(ctx)
	return err
}

type googleSubscription struct {
	cancel context.CancelFunc
}

func (g*googleSubscription) Close() error {
	g.cancel()
	return nil
}

func (g*googleProvider) getTopic(ctx context.Context, topicID string) (*pubsub.Topic, error) {
	topic := g.client.Topic(topicID)
	found, err := topic.Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !found {
		if !g.createTopic {
			return nil, fmt.Errorf("missing topic %s", topicID)
		}

		topic, err = g.client.CreateTopic(ctx, topicID)
		if err != nil {
			return nil, err
		}

	}

	return topic, nil
}

func (g*googleProvider) getSubscription(ctx context.Context, subscriptionName string, topic *pubsub.Topic) (*pubsub.Subscription, error) {
	sub := g.client.Subscription(subscriptionName)

	found, err := sub.Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !found {
		if !g.createSubscription {
			return nil, fmt.Errorf("unable to find subscription %s", subscriptionName)
		}

		sub, err = g.client.CreateSubscription(ctx, subscriptionName,
			pubsub.SubscriptionConfig{Topic: topic})

		if err != nil {
			return nil, err
		}
	}

	return sub, nil
}

func (g*googleProvider) Subscribe(ctx context.Context, topicID string, handler MsgHandler) (ISubscription, error) {
	topic, err := g.getTopic(ctx, topicID)
	if err != nil {
		return nil, err
	}

	subscriptionName := g.subscriptionSuffix + "-" + topicID
	sub, err := g.getSubscription(ctx, subscriptionName, topic)
	if err != nil {
		return nil, err
	}

	newContext, cancel := context.WithCancel(ctx)
	err = sub.Receive(newContext, func(ctx context.Context, m *pubsub.Message) {
		handler(m.Data)
		m.Ack()
	})
	if err != nil {
		defer cancel()
		return nil, err
	}

	return &googleSubscription{cancel: cancel}, nil
}


