package routes

import (
	"net/http"
	"sync"

	"github.com/jasonlvhit/gocron"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/config/client"
	"github.com/siddhant-vij/PokeChat-Universe/database"
	"github.com/siddhant-vij/PokeChat-Universe/routes/test/crud"
	"github.com/siddhant-vij/PokeChat-Universe/routes/test/health"
)

var (
	appConfig    *config.AppConfig
	dbService    *config.DbService
	redisService *config.RedisService
)

func init() {
	appConfig = &config.AppConfig{}
	config.LoadEnv(appConfig)

	appConfig.Mutex = &sync.RWMutex{}

	dbService = config.NewDatabaseService(appConfig)
	appConfig.DBQueries = database.New(dbService.DatabaseClient)

	redisService = config.NewRedisService(appConfig)
	appConfig.RedisClient = redisService.RedisClient

	client.FetchAndInsertRequest(appConfig)
}

func updateDatabaseCronJob() {
	gocron.Every(30).Days().Do(client.FetchAndInsertRequest, appConfig)
	<-gocron.Start()
}

func RegisterRoutes(mux *http.ServeMux) {
	// Cron job to update database
	go updateDatabaseCronJob()

	// Handlers for services setup, connections & CRUD operations
	HealthHandlers(mux)
	CrudHandlers(mux)
}

func HealthHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/health", health.ServerHealthHandler)

	mux.HandleFunc("/dbHealth", func(w http.ResponseWriter, r *http.Request) {
		health.DatabaseConnectionHealthHandler(w, r, dbService)
	})

	mux.HandleFunc("/redisHealth", func(w http.ResponseWriter, r *http.Request) {
		health.RedisConnectionHealthHandler(w, r, redisService)
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
