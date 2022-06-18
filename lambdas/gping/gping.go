package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Msg string `json:"msg"`
}

type Event struct {
	Input string `json:"input"`
}

func HandleRequest(event Event) (*Response, error) {
	log.Println("jeff:", event.Input)
	if len(event.Input) > 0 && event.Input == "ping" {
		return &Response{Msg: "pong"}, nil
	}
	return &Response{Msg: "I don't know"}, nil
}

func main() {

	lambda.Start(HandleRequest)
}
