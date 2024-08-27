package config

import (
	"sync"

	"github.com/redis/go-redis/v9"
	"golang.org/x/oauth2"

	"github.com/siddhant-vij/PokeChat-Universe/database"
)

type AppConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	RedisPort     string
	RedisAddress  string
	RedisDatabase string
	RedisPassword string

	DBQueries   *database.Queries
	RedisClient *redis.Client

	Mutex *sync.RWMutex

	AuthDomain       string
	ClientID         string
	ClientSecret     string
	CallbackURL      string
	SessionState     string
	PkceCodeVerifier string
	SessionTokenMap  map[string]*oauth2.Token
}
