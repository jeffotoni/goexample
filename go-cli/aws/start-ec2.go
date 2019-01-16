// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	//svc := ec2.New(session.New())
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")))

	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String("i-" + os.Getenv("AWS_INSTANCIE")),
		},
	}

	result, err := svc.StartInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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
