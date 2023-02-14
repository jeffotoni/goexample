// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// define flags
var (
	flagRegion   = flag.String("region", "", "example: us-east-1")
	flagInstance = flag.String("instance", "", "example: i-05bef719c14d68d20")
)

// init
func init() {

	flag.Parse()
	if *flagRegion == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *flagInstance == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func main() {

	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(*flagRegion))

	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(*flagInstance),
		},
	}
	// stop ec2
	result, err := svc.StopInstances(input)
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
