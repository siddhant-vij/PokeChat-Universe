package config

import (
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex"
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

	DBQueries   *pokedex.Queries
	RedisClient *redis.Client

	AuthDomain       string
	ClientID         string
	ClientSecret     string
	CallbackURL      string
	SessionState     string
	PkceCodeVerifier string

	IpAddress           string
	UserAgent           string
	AccessTokenIssuedAt time.Time
	AuthStatus          bool
	LoggedInUserId      uuid.UUID
	LoggedInUserName    string
}
