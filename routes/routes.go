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
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex"
	"github.com/siddhant-vij/PokeChat-Universe/middlewares"
	authroutes "github.com/siddhant-vij/PokeChat-Universe/routes/auth"
	chatroutes "github.com/siddhant-vij/PokeChat-Universe/routes/chat"
	pokedexroutes "github.com/siddhant-vij/PokeChat-Universe/routes/pokedex"
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
	appConfig.DBQueries = pokedex.New(dbService.DatabaseClient)

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

	// UI Handlers for Templ & Tailwind setup
	UiTestHandlers(mux)

	// Handlers for authentication
	AuthHandlers(mux)

	// Handlers for Home & Resource Pages
	PageHandlers(mux)

	// Handlers for App Workflow - Pokedex
	PokedexHandlers(mux)
	LoadMoreHandlers(mux)
	SearchAndSortHandlers(mux)
	ChatMessageHandlers(mux)
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

func UiTestHandlers(mux *http.ServeMux) {
	mux.Handle("/web", templ.Handler(test.Base()))
	mux.HandleFunc("/hello", ui.HelloWebHandler)
}

func AuthHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		authroutes.ServeLoginPage(w, r, authService, appConfig)
	})

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		authroutes.ServeCallbackPage(w, r, authService, appConfig)
	})

	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		authroutes.HandleLogout(w, r, appConfig)
	})
}

func PageHandlers(mux *http.ServeMux) {
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.ServeHomePage(w, r, appConfig)
	}))

	mux.Handle("/pokedex", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.ServeAvailablePage(w, r, appConfig)
	}), appConfig))

	mux.Handle("/collectedPokedex", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.ServeCollectedPage(w, r, appConfig)
	}), appConfig))

	mux.HandleFunc("/getPokemon", pokedexroutes.GetPokemonHandler)

	mux.Handle("/{pokemonName}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.ServePokemonPage(w, r, appConfig)
	}))

	mux.Handle("/collect", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.CollectPokemonHandler(w, r, appConfig)
	}), appConfig))

	mux.Handle("/collectPokemon", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.CollectPokemonHandlerOnPokemonPage(w, r, appConfig)
	}), appConfig))

	mux.Handle("/chat/", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chatroutes.PokedexChatHandler(w, r, appConfig)
	}), appConfig))

	mux.Handle("/add-pokemon", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chatroutes.AddPokemonToPokedexFromChatWindow(w, r, appConfig)
	}), appConfig))

	mux.Handle("/remove-pokemon", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chatroutes.RemovePokemonFromPokedexInChatWindow(w, r, appConfig)
	}), appConfig))
}

func PokedexHandlers(mux *http.ServeMux) {
	mux.Handle("/available", middlewares.IsAuthenticated(http.HandlerFunc(AvailableRedirectHandler), appConfig))

	mux.Handle("/collected", middlewares.IsAuthenticated(http.HandlerFunc(CollectedRedirectHandler), appConfig))

	mux.Handle("/pokeChat", middlewares.IsAuthenticated(http.HandlerFunc(ChatRedirectHandler), appConfig))
}

func LoadMoreHandlers(mux *http.ServeMux) {
	mux.Handle("/home-load-more", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.HomeAvailableLoadMore(w, r, appConfig)
	}))

	mux.Handle("/pa-load-more", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.PokedexAvailableLoadMore(w, r, appConfig)
	}), appConfig))

	mux.Handle("/pc-load-more", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.PokedexCollectedLoadMore(w, r, appConfig)
	}), appConfig))

	mux.Handle("/chat-load-more", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chatroutes.PokedexChatLoadMoreHandler(w, r, appConfig)
	}), appConfig))
}

func SearchAndSortHandlers(mux *http.ServeMux) {
	mux.Handle("/home-search", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.HomeAvailableSearch(w, r, appConfig)
	}))

	mux.Handle("/pa-search", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.PokedexAvailableSearch(w, r, appConfig)
	}), appConfig))

	mux.Handle("/pc-search", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.PokedexCollectedSearch(w, r, appConfig)
	}), appConfig))

	mux.Handle("/home-sort", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.HomeAvailableSort(w, r, appConfig)
	}))

	mux.Handle("/pa-sort", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.PokedexAvailableSort(w, r, appConfig)
	}), appConfig))

	mux.Handle("/pc-sort", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokedexroutes.PokedexCollectedSort(w, r, appConfig)
	}), appConfig))
}

func ChatMessageHandlers(mux *http.ServeMux) {
	mux.Handle("/chatMsg", middlewares.IsAuthenticated(http.HandlerFunc(chatroutes.ChatMessageHandler), appConfig))

	mux.Handle("/chatMsgBtn", middlewares.IsAuthenticated(http.HandlerFunc(chatroutes.ChatMessageButtonHandler), appConfig))
}
