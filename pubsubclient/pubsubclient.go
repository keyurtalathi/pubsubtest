package pubsubclient

import (
	"context"
	"time"

	"bitbucket.org/agrostar/onlineecomgateway/utils/loggerutils"
	"cloud.google.com/go/pubsub"
)

var logger = loggerutils.GetLogger()

//CreatePubSubClient a
func CreatePubSubClient(ctx context.Context, projectID string) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		logger.Errorf("\n\nERROR In PUBSUB client Create Client:\n\n", err)
		return nil, err
	}
	return client, nil
}

//CreateSubscription a
func CreateSubscription(ctx context.Context, client *pubsub.Client, topic *pubsub.Topic, subscriptionID string) (*pubsub.Subscription, error) {
	sub, err := client.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
		Topic:            topic,
		AckDeadline:      10 * time.Second,
		ExpirationPolicy: time.Duration(0),
	})
	if err != nil {
		logger.Errorf("\n\nERROR In PUBSUB client Create Subscription:\n\n", err)
		return nil, err
	}
	return sub, nil
}

//GetTopic a
func GetTopic(client *pubsub.Client, topicID string) *pubsub.Topic {
	return client.Topic(topicID)
}

//GetSubscription a
func GetSubscription(client *pubsub.Client, id string) *pubsub.Subscription {
	sub := client.Subscription(id)
	sub.ReceiveSettings.Synchronous = true
	sub.ReceiveSettings.MaxExtension = 0

	return sub
}

var data string

//PublishMessage a
func PublishMessage(ctx context.Context, topic *pubsub.Topic, message string) error {
	r := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(message),
	})
	id, err := r.Get(ctx)
	if err != nil {
		logger.Errorf("\n\nERROR In PUBSUB Publish Message\n\n", err)
		return err
	}
	logger.Info("Published a message with a message ID: " + id + ", And message: " + message)
	return nil
}
