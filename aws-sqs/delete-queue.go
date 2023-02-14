package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
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
		fmt.Println("Error Session SQS queue %v.", err)
		return
	}

	// Create a SQS service client.
	svc := sqs.New(sess)

	params := &sqs.DeleteQueueInput{
		QueueUrl: aws.String("https://sqs.us-east-1.amazonaws.com/xxxx/fwa1"),
	}

	result, err := svc.DeleteQueue(params)
	if err != nil {
		fmt.Println("Error  DeleteQueue: ", err)
		return
	}

	fmt.Println("Fila removida!")
	fmt.Println(result)
}
