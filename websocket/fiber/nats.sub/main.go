package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/nats-io/nats.go"
)

func main() {

	SubscribeAsync()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}

func SubscribeAsync() {

	//start := time.Now()
	chanpg := make(chan string)

	// Create server connection
	nc, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	// Subscribe to subject
	log.Printf("Subscribing to subject 'updates'\n")
	defer nc.Close()

	go func() {
		// Subscribe
		if _, err := nc.Subscribe("buy", func(msg *nats.Msg) {
			chanpg <- string(msg.Data)
		}); err != nil {
			log.Fatal(err)
		}

	}()

	// close(chanpg)
	//go func() {
	for {
		select {
		case cmsgJson := <-chanpg:
			// redis.SaveRedis(natsCount, cmsgJson)
			println("save:", cmsgJson)
		}
	}
}
