package routes

import (
	"log"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/database"
	"github.com/siddhant-vij/PokeChat-Universe/routes/test/crud"
	"github.com/siddhant-vij/PokeChat-Universe/routes/test/health"
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
	// Handlers for services setup, connections & CRUD operations
	HealthHandlers(mux)
	CrudHandlers(mux)
}

func HealthHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/health", health.ServerHealthHandler)

	mux.HandleFunc("/dbHealth", func(w http.ResponseWriter, r *http.Request) {
		health.ServiceConnectionHealthHandler(w, r, dbService)
	})

	mux.HandleFunc("/redisHealth", func(w http.ResponseWriter, r *http.Request) {
		health.ServiceConnectionHealthHandler(w, r, redisService)
	})
}

func CrudHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/dbCreate", func(w http.ResponseWriter, r *http.Request) {
		crud.DbCreateHandler(w, r, appConfig)
	})

	mux.HandleFunc("/dbRead", func(w http.ResponseWriter, r *http.Request) {
		crud.DbReadHandler(w, r, appConfig)
	})

	mux.HandleFunc("/dbUpdate", func(w http.ResponseWriter, r *http.Request) {
		crud.DbUpdateHandler(w, r, appConfig)
	})

	mux.HandleFunc("/dbDelete", func(w http.ResponseWriter, r *http.Request) {
		crud.DbDeleteHandler(w, r, appConfig)
	})

	mux.HandleFunc("/redisCreate", func(w http.ResponseWriter, r *http.Request) {
		crud.RedisCreateHandler(w, r, appConfig)
	})

	mux.HandleFunc("/redisRead", func(w http.ResponseWriter, r *http.Request) {
		crud.RedisReadHandler(w, r, appConfig)
	})

	mux.HandleFunc("/redisUpdate", func(w http.ResponseWriter, r *http.Request) {
		crud.RedisUpdateHandler(w, r, appConfig)
	})

	mux.HandleFunc("/redisDelete", func(w http.ResponseWriter, r *http.Request) {
		crud.RedisDeleteHandler(w, r, appConfig)
	})
}
