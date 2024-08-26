package services

import (
	"github.com/siddhant-vij/PokeChat-Universe/config"
)

type Service interface {
	Health() map[string]string
	Close(cfg *config.AppConfig) error
}
