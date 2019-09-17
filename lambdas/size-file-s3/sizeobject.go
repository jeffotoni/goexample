package main

import (
    "context"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "log"
)

var SIZE = int64(70000) // 7kb

type MsgResponse struct {
    Message string `json:"Answer:"`
}

func handlerS3(ctx context.Context, s3Event events.S3Event) (MsgResponse, error) {

    var msg string = "Events S3."
    for _, record := range s3Event.Records {

        rs3 := record.S3
        log.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, rs3.Bucket.Name, rs3.Object.Key)

        input := &s3.GetObjectInput{
            Bucket: aws.String(rs3.Bucket.Name),
            Key:    aws.String(rs3.Object.Key),
        }

        svc := s3.New(session.New())
        result, err := svc.GetObject(input)
        if err != nil {
            if aerr, ok := err.(awserr.Error); ok {
                switch aerr.Code() {
                case s3.ErrCodeNoSuchKey:
                    log.Println(s3.ErrCodeNoSuchKey, aerr.Error())
                default:
                    log.Println(aerr.Error())
                }
            } else {
                log.Println(err.Error())
            }
            continue
        }

        // == 7kb
        if aws.Int64Value(result.ContentLength) > SIZE {
            _, err := svc.DeleteObject(&s3.DeleteObjectInput{
                Bucket: aws.String(rs3.Bucket.Name),
                Key:    aws.String(rs3.Object.Key)})
            if err != nil {
                log.Printf("\nNão é possível excluir a chave %q do bucket %q, %v", rs3.Object.Key, rs3.Bucket.Name, err)
                continue
            }

            log.Println("Arquivo maior que o permitido: key: ", rs3.Object.Key, "excluido com [sucesso]")

        }
    }

    return MsgResponse{Message: msg}, nil
}

func main() {
    lambda.Start(handlerS3)
}
