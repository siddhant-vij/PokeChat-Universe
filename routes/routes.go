package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/database"
	"github.com/siddhant-vij/PokeChat-Universe/services"
	"github.com/siddhant-vij/PokeChat-Universe/services/postgres"
	"github.com/siddhant-vij/PokeChat-Universe/services/redis"
)

var (
	appConfig    *config.AppConfig
	dbService    services.Service
	redisService services.Service
)

func init() {
	appConfig = &config.AppConfig{}
	config.LoadEnv(appConfig)

	dbService = postgres.New(appConfig)
	tDB, ok := dbService.(*postgres.Service)
	if !ok {
		log.Fatalf("expected dbService to be postgres.Service, got %T", dbService)
	}
	appConfig.DBQueries = database.New(tDB.DB)

	redisService = redis.New(appConfig)
	rDB, ok := redisService.(*redis.Service)
	if !ok {
		log.Fatalf("expected redisService to be redis.Service, got %T", redisService)
	}
	appConfig.RedisClient = rDB.Redis
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", HealthHandler)

	mux.HandleFunc("/dbHealth", func(w http.ResponseWriter, r *http.Request) {
		ServiceConnectionHealthHandler(w, r, dbService)
	})

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
