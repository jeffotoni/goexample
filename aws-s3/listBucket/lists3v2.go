// Go in Acation
// @jeffotoni
// 2019-01-23

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// define flags
var (
	flagRegion = flag.String("region", "us-east-1", "example: eu-west-1")
	flagBucket = flag.String("bucket", "", "example: Mybucket")
	i          = 0
)

// init
func init() {

	flag.Parse()
	if *flagRegion == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *flagBucket == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func main() {

	// session aws, keys and secret and ~/.aws/credentials
	svc := s3.New(session.New(), &aws.Config{Region: aws.String(*flagRegion)})
	params := &s3.ListObjectsInput{
		Bucket: aws.String(*flagBucket),
	}

	// initializing line file
	var line string

	// list objects not bucket
	resp, _ := svc.ListObjects(params)
	for _, key := range resp.Contents {

		if *key.Size > 0 {
			i++
			line = strconv.Itoa(i) + "   " + *key.Key + "   " + strconv.FormatInt(*key.Size, 10)
			fmt.Println(line)
		}
	}
}
