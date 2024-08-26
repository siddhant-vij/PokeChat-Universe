package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

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
