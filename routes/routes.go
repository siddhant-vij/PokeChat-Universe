package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/jasonlvhit/gocron"
	"golang.org/x/oauth2"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/test"
	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/config/client"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/auth"
	"github.com/siddhant-vij/PokeChat-Universe/database"
	"github.com/siddhant-vij/PokeChat-Universe/middlewares"
	authroutes "github.com/siddhant-vij/PokeChat-Universe/routes/auth"
	"github.com/siddhant-vij/PokeChat-Universe/routes/test/crud"
	"github.com/siddhant-vij/PokeChat-Universe/routes/test/health"
	"github.com/siddhant-vij/PokeChat-Universe/routes/test/ui"
)

var (
	appConfig    *config.AppConfig
	dbService    *config.DbService
	redisService *config.RedisService
	authService  *auth.Authenticator
)

func init() {
	appConfig = &config.AppConfig{}
	config.LoadEnv(appConfig)

	appConfig.PkceCodeVerifier = oauth2.GenerateVerifier()
	appConfig.AuthStatus = false

	dbService = config.NewDatabaseService(appConfig)
	appConfig.DBQueries = database.New(dbService.DatabaseClient)

	redisService = config.NewRedisService(appConfig)
	appConfig.RedisClient = redisService.RedisClient

	client.FetchAndInsertRequest(appConfig)

	authService = auth.NewAuthenticator(appConfig)
}

func updateDatabaseCronJob() {
	gocron.Every(30).Days().Do(client.FetchAndInsertRequest, appConfig)
	<-gocron.Start()
}

func RegisterRoutes(mux *http.ServeMux) {
	// Cron job to update database
	go updateDatabaseCronJob()

	// File Server setup
	fileServer := http.FileServer(http.Dir("cmd/web/public"))
	mux.Handle("/cmd/web/public/", http.StripPrefix("/cmd/web/public/", fileServer))

	// Handlers for services setup, connections & CRUD operations
	HealthHandlers(mux)
	CrudHandlers(mux)

	// Handlers for authentication
	AuthHandlers(mux)

	// UI Handlers for Templ & Tailwind setup
	UiTestHandlers(mux)
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

func AuthHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		<body>
			<p>Home Page | <a href="/login">Login</a></p>			
		</body>
		</html>`))
	}))

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		authroutes.ServeLoginPage(w, r, authService, appConfig)
	})

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		authroutes.ServeCallbackPage(w, r, authService, appConfig)
	})

	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		authroutes.HandleLogout(w, r, appConfig)
	})

	mux.Handle("/resource", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		<body>
			<p>Resource Page | <a href="/logout">Logout</a></p>
		</body>
		</html>`))
	}), appConfig))
}

func UiTestHandlers(mux *http.ServeMux) {
	mux.Handle("/web", templ.Handler(test.Base()))
	mux.HandleFunc("/hello", ui.HelloWebHandler)
}
