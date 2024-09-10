package config

import (
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

var OpenAIClient *openai.Client

func InitOpenAI() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalf("OpenAI API key is not set")
	}

	OpenAIClient = openai.NewClient(apiKey)
}
