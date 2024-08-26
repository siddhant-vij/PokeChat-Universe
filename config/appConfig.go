package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
}

func LoadEnv(config *AppConfig) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file, Err: %v", err)
	}

	config.DBHost = os.Getenv("DB_HOST")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBUser = os.Getenv("DB_USERNAME")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_DATABASE")

	config.RedisPort = os.Getenv("REDIS_PORT")
	config.RedisAddress = os.Getenv("REDIS_ADDRESS")
	config.RedisDatabase = os.Getenv("REDIS_DATABASE")
	config.RedisPassword = os.Getenv("REDIS_PASSWORD")
}
