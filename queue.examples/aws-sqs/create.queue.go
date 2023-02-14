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

	policy := `{
	  "Version": "2012-10-17",
	  "Id": "arn:aws:sqs:us-east-1:xxxxxxx:nameQueue/SQSDefaultPolicy",
	  "Statement": [
	    {
	      "Sid": "Sid1567608179743",
	      "Effect": "Allow",
	      "Principal": "*",
	      "Action": "SQS:*",
	      "Resource": "arn:aws:sqs:us-east-1:xxxxxxx:nameQueue"
	    }
	  ]
	}`

	params := &sqs.CreateQueueInput{
		QueueName: aws.String("nameQueue"), // Required
		Attributes: map[string]*string{
			//"Key":                           aws.String("nameQueue"), // Required
			"MessageRetentionPeriod":        aws.String("14400"), // retencao na fila, 4 horas.
			"ReceiveMessageWaitTimeSeconds": aws.String("20"),    // 20 seconds para aguardar antes de entregar a msgs
			"VisibilityTimeout":             aws.String("40"),
			"DelaySeconds":                  aws.String("10"),   // Atraso na entrega
			"Policy":                        aws.String(policy), // Atraso na entrega
			// More values...
		},
	}

	result, err := svc.CreateQueue(params)
	if err != nil {
		fmt.Println("Error createQueue: ", err)
		return
	}

	fmt.Println("Fila criada com sucesso!")
	fmt.Println(result)
}
