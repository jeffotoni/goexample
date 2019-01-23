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

var (
	pathCsv    = "./list-bucket.csv"
	flagRegion = flag.String("region", "us-east-1", "example:")
	flagBucket = flag.String("bucket", "myBucket", "example:")
)

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

	svc := s3.New(session.New(), &aws.Config{Region: aws.String(*flagRegion)})
	params := &s3.ListObjectsInput{
		Bucket: aws.String(*flagBucket),
		//Prefix: aws.String("root/subfolder"),
	}

	// remove path
	os.Remove(pathCsv)

	// open and new file append
	f, err := os.OpenFile(pathCsv, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer f.Close()

	var line string
	line = "line,object name,size\n"

	if _, err = f.WriteString(line); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Generate file: list-bucket.csv of " + *flagBucket + "")
	var i int
	resp, _ := svc.ListObjects(params)
	for _, key := range resp.Contents {
		if *key.Size > 0 {
			i++
			line = strconv.Itoa(i) + "," + *key.Key + "," + strconv.FormatInt(*key.Size, 10) + "\n"
			f.WriteString(line)
		}
	}

	fmt.Println("file generated successfully.")
}
