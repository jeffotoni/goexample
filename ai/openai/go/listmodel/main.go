package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Substitua YOUR_OPENAI_API_KEY pela sua chave de API da OpenAI
	apiKey := "OPENAI_API_KEY"
	client := openai.NewClient(os.Getenv(apiKey))

	// Lista os modelos disponíveis
	models, err := client.ListModels(context.Background())
	if err != nil {
		log.Fatalf("Failed to list models: %v", err)
	}

	// Imprime os modelos disponíveis
	for _, model := range models.Models {
		fmt.Printf("Model ID: %s - Name: %s\n", model.ID, model.Object)
	}
}
