package main

import (
	"context"
	"github.com/GeniusAI-Platform/openai"
	"github.com/GeniusAI-Platform/openai/client"
	"github.com/GeniusAI-Platform/openai/patterns/completion"
	"github.com/GeniusAI-Platform/openai/types/programming"
	"github.com/GeniusAI-Platform/openai/models"
	"log"
	"os"
)

var code string = `
func add(a, b int) int {
	return a + b
}
`

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	cli, err := client.New([]string{apiKey})
	if err != nil {
		log.Fatalln(err)
	}

	c := openai.NewCompletion(cli)
	resp, err := c.CreateCompletionFromPattern(context.Background(), completion.ProgrammingLanguageTranslator(
		Model:  models.TEXT_DAVINCI_002,
		code,
		programming.Go,
		programming.Python,
		0,
	))

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Choices[0].Text)
}

