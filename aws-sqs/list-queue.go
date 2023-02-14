package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"time"
)

// Usage:
// go run sqs_listqueues.go
func main() {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		fmt.Println("Error list SQS queue %v.", err)
	}

	// Create a SQS service client.
	svc := sqs.New(sess)

	go func() {

		for {
			// List the queues available in a given region.
			result, err := svc.ListQueues(nil)
			if err != nil {
				fmt.Println("Error", err)
				return
			}

			fmt.Println("Success")
			// As these are pointers, printing them out directly would not be useful.
			for i, urls := range result.QueueUrls {
				// Avoid dereferencing a nil pointer.
				if urls == nil {
					continue
				}
				time.Sleep(time.Millisecond * 600)
				fmt.Printf("%d: %s\n", i, *urls)
			}

			time.Sleep(time.Second * 3)
		}
	}()

	time.Sleep(time.Hour * 1)
}
