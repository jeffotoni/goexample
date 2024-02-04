package main

import (
	"context"
	"log"
	"os"

	openai "github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/completion"
)

func main() {
	ctx := context.Background()
	s := openai.NewSession(os.Getenv("OPENAI_API_KEY"))

	temperature := 0.7

	client := completion.NewClient(s, "gpt-3.5-turbo-0301")
	resp, err := client.Create(ctx, &completion.CreateParams{
		N:           1,
		MaxTokens:   200,
		Temperature: temperature,
		Prompt:      []string{"Qual é a capital da França?"},
	})
	if err != nil {
		log.Fatalf("Failed to complete: %v", err)
	}

	for _, choice := range resp.Choices {
		log.Println(choice.Text)
	}
}
