/*
* Golang notify
*
* @package     main
* @author      @jeffotoni
* @size        2018
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Receive message from Queue with long polling enabled.
//
// Usage:
//    go run sqs_longpolling_receive_message.go -n queue_name -t timeout
func main() {

	var name string
	var timeout int64

	flag.StringVar(&name, "n", "", "Queue name")
	flag.Int64Var(&timeout, "t", 20, "(Optional) Timeout in seconds for long polling")
	flag.Parse()

	if len(name) == 0 {
		flag.PrintDefaults()
		exitErrorf("Queue name required")
	}

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create a SQS service client.
	svc := sqs.New(sess)

	// Need to convert the queue name into a URL. Make the GetQueueUrl
	// API call to retrieve the URL. This is needed for receiving messages
	// from the queue.
	resultURL, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == sqs.ErrCodeQueueDoesNotExist {
			exitErrorf("Error Unable to find queue::: %q.", name)
		}
		exitErrorf("Unable to queue %q, %v.", name, err)
	}

	go func(err error) {

		for {

			//fmt.Println("::: I :::", i)

			time.Sleep(time.Second * 2)

			// Receive a message from the SQS queue with long polling enabled.
			result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
				QueueUrl: resultURL.QueueUrl,
				AttributeNames: aws.StringSlice([]string{
					"SentTimestamp",
				}),

				MaxNumberOfMessages: aws.Int64(1),
				MessageAttributeNames: aws.StringSlice([]string{
					"All",
				}),

				WaitTimeSeconds: aws.Int64(timeout),
			})

			if err != nil {
				exitErrorf("Unable to receive message from queue %q, %v.", name, err)
			}

			fmt.Printf("Received %d messages.\n", len(result.Messages))

			if len(result.Messages) > 0 {

				fmt.Println(result.Messages)

				fmt.Println("aguardando 10s para remover")
				time.Sleep(time.Second * 10)

				qURL := os.Getenv("qURL")

				resultDelete, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
					QueueUrl:      &qURL,
					ReceiptHandle: result.Messages[0].ReceiptHandle,
				})

				if err != nil {
					fmt.Println("Delete Error", err)
					return
				}

				fmt.Println("Message Deleted", resultDelete)
			}

		}
	}(err)

	time.Sleep(time.Hour * 1)

}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
