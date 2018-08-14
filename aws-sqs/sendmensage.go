/*
* Golang notify
*
* @package     main
* @author      @jeffotoni
* @size        2018
 */

package main

import (
	"fmt"
	"os"
	"time"
)

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	go func() {

		// export qURL=http://colocaaquisuaurl
		// URL to our queue
		qURL := os.Getenv("qURL")

		i := 0

		for {

			result, err := svc.SendMessage(&sqs.SendMessageInput{
				// somente para FILAS FIFO
				// MessageGroupId: aws.String("grupo_fifo_" + fmt.Sprintf("%d", i)),

				DelaySeconds: aws.Int64(10),
				MessageAttributes: map[string]*sqs.MessageAttributeValue{
					"Title": {
						DataType:    aws.String("String"),
						StringValue: aws.String("Golang AWS SQS 10" + fmt.Sprintf("%d", i)),
					},
					"Author": {
						DataType:    aws.String("String"),
						StringValue: aws.String("JeffOtoni"),
					},
					"WeeksOn": {
						DataType:    aws.String("Number"),
						StringValue: aws.String("1000"),
					},
				},

				MessageBody: aws.String("Information about current NY Times fiction bestseller for week of 12/08/2018."),
				QueueUrl:    &qURL,
			})

			if err != nil {
				fmt.Println("Error", err)
				return
			}

			fmt.Println("Success", *result.MessageId)
			i++
			time.Sleep(time.Second * 2)
		}
	}()

	time.Sleep(time.Hour * 1)
}
