package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/services"
	"github.com/siddhant-vij/PokeChat-Universe/services/postgres"
	"github.com/siddhant-vij/PokeChat-Universe/services/redis"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", HealthHandler)

	var dbService = postgres.New()
	mux.HandleFunc("/dbHealth", func(w http.ResponseWriter, r *http.Request) {
		ServiceConnectionHealthHandler(w, r, dbService)
	})

	var redisService = redis.New()
	mux.HandleFunc("/redisHealth", func(w http.ResponseWriter, r *http.Request) {
		ServiceConnectionHealthHandler(w, r, redisService)
	})
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func ServiceConnectionHealthHandler(w http.ResponseWriter, r *http.Request, s services.Service) {
	jsonResp, err := json.Marshal(s.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
