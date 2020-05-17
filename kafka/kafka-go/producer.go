package main

import (
	"context"
	"fmt"
	//"os"
	"flag"
	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
	"strconv"
	"time"
)

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:        []string{kafkaURL},
		Topic:          topic,
		Balancer:       &kafka.LeastBytes{},
		RequiredAcks:   1,           // acks = 0, acks = 1, acks = -1
		CommitInterval: time.Second, // flushes commits to Kafka every second
	})
}

func main() {

	flagHost := flag.String("host", "localhost:9092", "string")
	flagTopic := flag.String("topic", "topicgo1", "string")
	flag.Parse()

	// get kafka writer using environment variables.
	kafkaURL := *flagHost
	topic := *flagTopic

	fmt.Println("Url: ", kafkaURL)
	fmt.Println("Topic: ", topic)

	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()
	fmt.Println("Go Start Producing ... !!")
	for i := 0; ; i++ {
		uuid := fmt.Sprint(uuid.New())
		msgJson := `{"uuid":"` + uuid + `", "key":` + strconv.Itoa(i) + `,"msg":success", "event":"kafka test"}`
		msg := kafka.Message{
			Key:   []byte(uuid),
			Value: []byte(msgJson),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Key-", i)
		time.Sleep(time.Millisecond * 300)
	}
}
