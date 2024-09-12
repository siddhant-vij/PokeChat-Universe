package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ollama/ollama/api"
)

type OllamaService struct {
	Client *api.Client
}

var ollamaServiceInstance *OllamaService

func NewOllamaService() *OllamaService {
	if ollamaServiceInstance != nil {
		return ollamaServiceInstance
	}

	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatalf("Error creating Ollama API client. Err: %v", err)
	}

	ollamaServiceInstance = &OllamaService{
		Client: client,
	}

	return ollamaServiceInstance
}

func (s *OllamaService) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Sending a simple request to check if the Ollama server is up
	err := s.Client.Heartbeat(ctx)
	if err != nil {
		log.Printf("Ollama server is down: %v", err)
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("Ollama server error: %v", err)
		return stats
	}

	// If the server responds, update the stats accordingly
	stats["status"] = "up"
	stats["message"] = "Ollama server is healthy"
	stats["ollama_version"], _ = s.Client.Version(ctx)

	return stats
}
