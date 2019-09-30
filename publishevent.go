package pubsubutils

import (
	"context"

	"bitbucket.org/agrostar/onlineecomgateway/utils/pubsubutils/pubsubclient"
)

func PublishEvent(ctx context.Context, projectID, topicID, data string) {
	client, _ := pubsubclient.CreatePubSubClient(ctx, projectID)
	topic := pubsubclient.GetTopic(client, topicID)
	defer topic.Stop()
	pubsubclient.PublishMessage(ctx, topic, data)
}
