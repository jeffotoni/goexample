// @package     Golang rsync Filework One
// @author      @jeffotoni

package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"
)

func Download() error {

	s3Config = &aws.Config{
		Credentials: credentials.NewStaticCredentials(AWS_KEY, AWS_SECRET, ""),
		Region:      aws.String("us-east-1"),
	}

	awsSession := session.New(s3Config)
	s3Svc := s3.New(awsSession)

	fooFile, err := os.OpenFile("/tmp/camiseta1.png", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error create, open file: ", err.Error())
		return err
	}

	result, err := s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("your-bucket"),
		Key:    aws.String("camiseta1.png"),
	})

	if err != nil {
		fmt.Println("Error getObject: ", err.Error())
		return err
	}

	n, err := io.Copy(fooFile, result.Body)
	result.Body.Close()

	fmt.Println("Teste Download file [sucesso].....", n)
	return nil
}
