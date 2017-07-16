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
	"crypto/md5"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"math/rand"
	// "os"
)

func main() {

	config := &aws.Config{
		Region: aws.String("us-east-1"),
	}

	svc := dynamodb.New(session.New(config))

	var login_name string
	var cont string
	var passRand string
	var data []byte
	var password string
	var seed string
	var resultString string

	for i := 10; i < 1100; i++ {

		login_name = "login_" + fmt.Sprintf("%d", i) + "@gmail.com"

		cont = fmt.Sprintf("%d", i)

		seed = fmt.Sprintf("%s", Seed(1938, 24509+i))

		passRand = login_name + cont + seed

		data = []byte(passRand)

		password = fmt.Sprintf("%x", md5.Sum(data))

		// fmt.Println(login_name)
		// os.Exit(1)

		input := &dynamodb.PutItemInput{
			Item: map[string]*dynamodb.AttributeValue{
				"id": {
					N: aws.String(cont),
				},
				"login": {
					S: aws.String(login_name),
				},
				"password": {
					S: aws.String(password),
				},
			},

			ReturnConsumedCapacity: aws.String("TOTAL"),
			TableName:              aws.String("ukklogin"),
		}

		result, err := svc.PutItem(input)

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeConditionalCheckFailedException:
					fmt.Println(dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
				case dynamodb.ErrCodeProvisionedThroughputExceededException:
					fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
				case dynamodb.ErrCodeResourceNotFoundException:
					fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
				case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
					fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
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

		resultString = fmt.Sprintf("%s", result)
	}

	fmt.Println(resultString)
}

func Seed(min int, max int) int {

	return (min + rand.Intn(max-min))
}
