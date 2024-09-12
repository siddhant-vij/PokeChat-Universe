package chat

import (
	"context"

	"github.com/ollama/ollama/api"
	
	"github.com/siddhant-vij/PokeChat-Universe/config"
)

func GenerateChatResponse(userMessage string, cfg *config.AppConfig) <-chan string {
	responseStream := make(chan string)

	go func() {
		defer close(responseStream)
		messages := []api.Message{
			{
				Role:    "user",
				Content: userMessage,
			},
		}
		req := &api.ChatRequest{
			Model:    "llama3.1",
			Messages: messages,
		}
		responseFunc := func(resp api.ChatResponse) error {
			responseStream <- resp.Message.Content
			return nil
		}
		cfg.OllamaClient.Chat(context.Background(), req, responseFunc)
	}()

	return responseStream
}
