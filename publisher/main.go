package main

import (
	"context"
	"os"
	"time"

	"bitbucket.org/agrostar/onlineecomgateway/utils/loggerutils"
	"bitbucket.org/agrostar/onlineecomgateway/utils/pubsubutils/pubsubclient"
)

var logger = loggerutils.GetLogger()

func main() {
	ctx := context.Background()
	projectID := os.Getenv("PROJECT")
	topicID := os.Getenv("TOPIC")

	client, _ := pubsubclient.CreatePubSubClient(ctx, projectID)
	// if err != nil {
	// 	logger.Info("IN MAIN")

	topic := pubsubclient.GetTopic(client, topicID)
	defer topic.Stop()

	for {
		pubsubclient.PublishMessage(ctx, topic, "my name is khan")
		time.Sleep(3 * time.Second)
	}
}
