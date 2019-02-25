// Go in action
// @jeffotoni
// 2019-01-16

// example:
// https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/ec2/create_image_no_block_device.go
package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"

    "fmt"
)

func main() {
    // Load session from shared config
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create EC2 service client
    svc := ec2.New(sess)

    opts := &ec2.CreateImageInput{
        Description: aws.String("image description"),
        InstanceId:  aws.String("i-abcdef12"),
        Name:        aws.String("image name"),
        BlockDeviceMappings: []*ec2.BlockDeviceMapping{
            {
                DeviceName: aws.String("/dev/sda1"),
                NoDevice:   aws.String(""),
            },
            {
                DeviceName: aws.String("/dev/sdb"),
                NoDevice:   aws.String(""),
            },
            {
                DeviceName: aws.String("/dev/sdc"),
                NoDevice:   aws.String(""),
            },
        },
    }
    resp, err := svc.CreateImage(opts)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("ID: ", resp.ImageId)
}
