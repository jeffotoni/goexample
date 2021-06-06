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

	for {
		t := time.Now().Format("2006-01-02 15:04:05")
	if err := nc.Publish("updates", []byte(`{"name":"jeffotoni_`+t+`","time":"`+t+`"}`)); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Millisecond * 100)
	}

	println("done..")
}
