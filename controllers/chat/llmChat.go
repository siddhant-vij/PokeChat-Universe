package chat

import (
	"context"
	"strings"

	"github.com/ollama/ollama/api"

	"github.com/siddhant-vij/PokeChat-Universe/config"
)

func GenerateChatResponse(userMessage, pokemonName string, cfg *config.AppConfig) <-chan string {
	responseStream := make(chan string)

	go func() {
		defer close(responseStream)
		messages := GenerateMessage(userMessage, pokemonName, cfg)
		req := &api.ChatRequest{
			Model:    "llama3.1",
			Messages: messages,
		}
		responseFunc := func(resp api.ChatResponse) error {
			content := resp.Message.Content
			content = strings.ReplaceAll(content, "\n", "<br />")
			// Scope to build a library for rendering streaming markdown into HTML (frontend)
			responseStream <- content
			return nil
		}
		cfg.OllamaClient.Chat(context.Background(), req, responseFunc)
	}()

	return responseStream
}
