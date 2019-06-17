package main

import (
	"fmt"
)

type Notification struct {
	Attachments []struct {
		Title    string   `json:"title"`
		Text     string   `json:"text"`
		MrkdwnIn []string `json:"mrkdwn_in"`
	} `json:"attachments"`
}

func main() {
	fmt.Println("Hello, playground")

	message := &Notification{
		Attachments: []struct {
			Title    string   `json:"title"`
			Text     string   `json:"text"`
			MrkdwnIn []string `json:"mrkdwn_in"`
		}{
			{Title: "Teste", Text: "sending from go", MrkdwnIn: []string{"text", "pretext"}},
		},
	}

	fmt.Println(message)
}
