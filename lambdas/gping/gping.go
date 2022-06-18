package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Msg string `json:"msg"`
}

type Event struct {
	Name string `json:"name"`
}

func HandleRequest(event Event) (*Response, error) {

	return &Response{Msg: "pong"}, nil
}

func main() {

	lambda.Start(HandleRequest)
}
