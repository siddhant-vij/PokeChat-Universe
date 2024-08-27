package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(cfg *AppConfig) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file, Err: %v", err)
	}

	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")
	cfg.DBUser = os.Getenv("DB_USERNAME")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBName = os.Getenv("DB_DATABASE")

	cfg.RedisPort = os.Getenv("REDIS_PORT")
	cfg.RedisAddress = os.Getenv("REDIS_ADDRESS")
	cfg.RedisDatabase = os.Getenv("REDIS_DATABASE")
	cfg.RedisPassword = os.Getenv("REDIS_PASSWORD")

	cfg.AuthDomain = os.Getenv("AUTH0_DOMAIN")
	cfg.ClientID = os.Getenv("AUTH0_CLIENT_ID")
	cfg.ClientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	cfg.CallbackURL = os.Getenv("AUTH0_CALLBACK_URL")
}
