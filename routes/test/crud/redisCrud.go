package crud

import (
	"context"
	"log"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/config"
)

var kvPair = map[string]string{
	"test": "value",
}

func RedisCreateHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	err := cfg.RedisClient.Set(context.Background(), "test", kvPair["test"], 0).Err()
	if err != nil {
		log.Fatalf("error inserting kvPair. Err: %v", err)
	}

	w.Write([]byte("kvPair inserted in redis!"))
}

func RedisReadHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	val, err := cfg.RedisClient.Get(context.Background(), "test").Result()
	if err != nil {
		log.Fatalf("error getting kvPair. Err: %v", err)
	}

	w.Write([]byte(val + " is obtained from redis!"))
}

func RedisUpdateHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	err := cfg.RedisClient.Set(context.Background(), "test", "valueUpdated", 0).Err()
	if err != nil {
		log.Fatalf("error updating kvPair. Err: %v", err)
	}

	w.Write([]byte("value updated in redis!"))
}

func RedisDeleteHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	err := cfg.RedisClient.Del(context.Background(), "test").Err()
	if err != nil {
		log.Fatalf("error deleting kvPair. Err: %v", err)
	}

	w.Write([]byte("kvPair deleted from redis!"))
}
