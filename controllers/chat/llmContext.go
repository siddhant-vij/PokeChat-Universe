package chat

import (
	"github.com/ollama/ollama/api"

	"github.com/siddhant-vij/PokeChat-Universe/config"
)

func GenerateMessage(currentUserMessage, pokemonName string, cfg *config.AppConfig) []api.Message {
	messages := []api.Message{
		{
			Role:    "user",
			Content: currentUserMessage,
		},
	}
	return messages
}
