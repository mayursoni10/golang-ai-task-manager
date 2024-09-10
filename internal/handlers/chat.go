package handlers

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayursoni10/golang-ai-task-manager/internal/config"
)

func ChatHandler(c *gin.Context) {
	var request struct {
		Message string `json:"message"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	response, err := config.OpenAIClient.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: request.Message,
			},
		},
	})

	if err != nil {
		log.Printf("Error from OpenAI API: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response from OpenAI"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response.Choices[0].Message.Content})
}
