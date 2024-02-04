package main

import (
	"context"
	"fmt"
	"log"
	"os"

	gogpt "github.com/bakks/go-gpt3"
)

func main() {
	// Substitua YOUR_OPENAI_API_KEY pela sua chave de API da OpenAI
	apiKey := os.Getenv("OPENAI_API_KEY")
	model := "gpt-3.5-turbo-1106"
	client := gogpt.NewClient(apiKey, gogpt.WithDefaultEngine(model))

	// client := gogpt.NewClient(apiKey, modelOption.WithDefaultEngine(""))
	// Configura os parâmetros da solicitação
	prompt := []string{"Qual é a capital da França?"}
	temperature := float32(0.7)
	maxTokens := 100

	// Prepara a solicitação de Completion, ajustando campos conforme necessário
	req := gogpt.CompletionRequest{
		Prompt:      prompt,
		Temperature: &temperature,
		MaxTokens:   &maxTokens,
	}

	// Envia a solicitação e recebe a resposta
	ctx := context.Background()
	resp, err := client.Completion(ctx, req) // Ajuste para o método correto
	if err != nil {
		log.Fatalf("Failed to create completion: %v", err)
	}

	// Imprime a resposta, ajuste conforme a estrutura da resposta
	fmt.Println("Resposta:", resp.Choices[0].Text)
}
