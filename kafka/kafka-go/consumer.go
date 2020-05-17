package main

import (
	"context"
	"fmt"
	"log"
	//"os"
	"flag"
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func main() {

	flagTopic := flag.String("topic", "topicgo1", "string")
	flagGroup := flag.String("group", "logger-group1", "string")
	flag.Parse()

	// get kafka reader using environment variables.
	kafkaURL := "localhost:9092" //os.Getenv("kafkaURL")
	topic := *flagTopic          //os.Getenv("topic")
	groupID := *flagGroup        //os.Getenv("groupID")

	fmt.Println("Url: ", kafkaURL)
	fmt.Println("Topic: ", topic)
	fmt.Println("Group: ", groupID)

	reader := getKafkaReader(kafkaURL, topic, groupID)

	defer reader.Close()

	ticker := time.NewTicker(900 * time.Millisecond)
	done := make(chan bool)
	go func() {
		fmt.Println("start consuming ... !!")
		for {

			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("consumer: ", t.Format("2006-01-02 15:04:05"))
				m, err := reader.ReadMessage(context.Background())
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
			}
		}
	}()

	time.Sleep(1000000 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
