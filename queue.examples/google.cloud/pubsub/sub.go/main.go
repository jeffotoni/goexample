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
	subscript = "pull"
)

func main() {
	subscriptp := flag.String("subscription", subscript, "string")
	flag.Parse()
	if subscriptp != nil {
		subscript = *subscriptp
	}

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		client, err := pubsub.NewClient(ctx, projectID)
		if err != nil {
			log.Printf("Error NewClient:%s", err.Error())
			return
		}

		sub := client.Subscription(subscript)
		err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
			fmt.Println(string(m.Data))
			m.Ack()
		})
		if err != nil {
			log.Println(err)
		}
		<-time.After(time.Second * 5)
	}
}
