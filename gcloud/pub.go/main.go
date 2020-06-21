package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
)

var (
	projectID = "projeto-eng1"
	nameTopic = "topic.pub.go"
	msg       = `{"ID":"39339xxx9w33","msg":"testando envio de mensagem"}`
)

func main() {
	msgp := flag.String("msg", msg, "string")
	flag.Parse()
	if msgp != nil {
		msg = *msgp
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Error NewClient:%s", err.Error())
		return
	}

	topic := client.Topic(nameTopic)
	defer topic.Stop()
	var results []*pubsub.PublishResult

	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(msg)},
	)

	results = append(results, res)
	// Do other work ...
	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			log.Printf("Error Get results:%s", err.Error())
			continue
		}
		fmt.Printf("Published a message with a message ID: [%s]\n", id)
	}
}
