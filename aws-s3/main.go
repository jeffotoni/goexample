package main

import (
  "log"
  "time"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3"
)

func main() {

  // awsKey := "xxxxxxxxxxxxxxxxxxxx"
  // awsSecret := "yyyyyyyyyyyyyyyyyyyyy"
  // region := "us-east-1"

  bucket := "fileslivecycle"
  filename := "super-jeff.png"

  sess := session.Must(session.NewSession())

  // creds := credentials.NewStaticCredentials(awsKey, awsSecret, "")
  // _, err := creds.Get()
  // if err != nil {
  //   panic(err)
  // }
  // cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds).WithLogLevel(aws.LogDebugWithHTTPBody)
  // sess := session.Must(session.NewSession(cfg))

  path := "/" + filename

  svc := s3.New(sess)
  req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(path),
  })

  url, err := req.Presign(2 * time.Minute)

  if err != nil {
    log.Fatalf("[ERROR] Failed to sign request.", err)
  }

  log.Println(url)
}
