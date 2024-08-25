package main

import (
	"log"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/middlewares"
	"github.com/siddhant-vij/PokeChat-Universe/routes"
)

func main() {
	mux := http.NewServeMux()
	corsMux := middlewares.CorsMiddleware(mux)
	routes.RegisterRoutes(mux)

	log.Fatal(http.ListenAndServe(":8080", corsMux))
}
