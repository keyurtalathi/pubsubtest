package main

import (
	"context"
	"fmt"
	"os"

	"bitbucket.org/agrostar/onlineecomgateway/utils/loggerutils"
	"bitbucket.org/agrostar/onlineecomgateway/utils/pubsubutils/pubsubclient"
	"cloud.google.com/go/pubsub"
)

var logger = loggerutils.GetLogger()

func xyz() {
	fmt.Println("a ...interface{}")
}
func main() {
	ctx := context.Background()
	projectID := os.Getenv("PROJECT")
	subscriptionID := os.Getenv("SUBSCRIPTION")

	client, _ := pubsubclient.CreatePubSubClient(ctx, projectID)
	// if err != nil {
	// 	logger.Info("IN MAIN")

	sub := pubsubclient.GetSubscription(client, subscriptionID)
	if err := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Println(string(m.Data))
		xyz()
		m.Ack()

	}); err != nil {
		logger.Info(err.Error())
	}

}
