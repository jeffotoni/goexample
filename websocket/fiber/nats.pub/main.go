package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	println(nats.DefaultURL)
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	//for {
	if err := nc.Publish("updates", []byte(`{"name":"jeffotoni","code":"x030303030303030"}`)); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Millisecond * 100)
	//}

	println("done..")
}
