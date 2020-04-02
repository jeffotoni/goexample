package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Path: ", req.Path)
	log.Println("All ", req)
	if req.Path == "/users" {
		if req.HTTPMethod == "GET" {
			return handleGetUsers(req)
		}
		if req.HTTPMethod == "POST" {
			return handleCreateUser(req)
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       http.StatusText(http.StatusMethodNotAllowed),
	}, nil
}

func handleGetUsers(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var users string = `[{"id":3939,"name":"jeffotoni", "status":"ativo"},{"id":100,"name":"Yara", "status":"ativo"}]`
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       users,
	}, nil
}

func handleCreateUser(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user User
	err := json.Unmarshal([]byte(req.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "Created",
	}, nil
}

func main() {
	lambda.Start(router)
}
