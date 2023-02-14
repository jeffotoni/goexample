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
	/*
			Addresses: [
		    {
		      AllocationId: "eipalloc-05ca438ace2e4a2d4",
		      AssociationId: "eipassoc-085577a46b88b61ba",
		      Domain: "vpc",
		      InstanceId: "i-0512874c36c8205a0",
		      NetworkInterfaceId: "eni-0494dcd054ec93cb6",
		      NetworkInterfaceOwnerId: "873761630739",
		      PrivateIpAddress: "172.30.0.248",
		      PublicIp: "18.215.29.129",
		      PublicIpv4Pool: "amazon",
		      Tags: [{
		          Key: "Name",
		          Value: "filserver/go"
		        }]
		    },
	*/

	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")))
	input := &ec2.DescribeAddressesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(""),
				Values: []*string{
					aws.String(""),
				},
			},
		},
	}

	result, err := svc.DescribeAddresses(input)
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
