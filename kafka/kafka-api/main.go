package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	kafka "github.com/segmentio/kafka-go"
)

func producerHandler(kafkaWriter *kafka.Writer) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Println("Body:", err)
				return
			}
			msg := kafka.Message{
				Key:   []byte(fmt.Sprintf("address-%s", req.RemoteAddr)),
				Value: body,
			}
			err = kafkaWriter.WriteMessages(req.Context(), msg)
			if err != nil {
				wrt.WriteHeader(200)
				wrt.Write([]byte(err.Error()))
				log.Println("kafka:", err)
			}
		} else {
			wrt.WriteHeader(400)
			wrt.Write([]byte(`{"msg":"Method not allowed"}`))
		}
	})
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
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

	kafkaWriter := getKafkaWriter(kafkaURL, topic)

	defer kafkaWriter.Close()

	// Add handle func for producer.
	http.HandleFunc("/producer", producerHandler(kafkaWriter))

	// Run the web server.
	fmt.Println("start producer-api 8181 ... !!")
	log.Fatal(http.ListenAndServe(":8181", nil))
}
