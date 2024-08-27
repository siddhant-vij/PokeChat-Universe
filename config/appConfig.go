package config

import (
	"sync"
	"time"

	"github.com/redis/go-redis/v9"

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

	IpAddress           string
	UserAgent           string
	AccessTokenIssuedAt time.Time
}
