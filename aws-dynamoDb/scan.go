/*
* Example tollbooth
*
* @package     main
* @author      @jeffotoni
* @size        16/07/2017
*
 */

package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {

	config := &aws.Config{
		Region: aws.String("us-east-1"),
	}

	svc := dynamodb.New(session.New(config))

	input := &dynamodb.ScanInput{

		// ExpressionAttributeNames: map[string]*string{

		// 	"AT": aws.String("id"),
		// 	"ST": aws.String("login"),
		// },
		// ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
		// 	":a": {
		// 		S: aws.String("No One You Know"),
		// 	},
		// },

		// FilterExpression:     aws.String("login = :jeff"),
		// ProjectionExpression: aws.String(""),
		TableName: aws.String("ukklogin"),
	}

	result, err := svc.Scan(input)

	if err != nil {

		if aerr, ok := err.(awserr.Error); ok {

			switch aerr.Code() {

			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())

			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())

			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())

			default:
				fmt.Println(aerr.Error())

			}

		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

}
