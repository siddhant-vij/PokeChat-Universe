package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/services"
	"github.com/siddhant-vij/PokeChat-Universe/services/postgres"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", HealthHandler)

	var s = postgres.New()
	mux.HandleFunc("/dbHealth", func(w http.ResponseWriter, r *http.Request) {
		DbConnectionHealthHandler(w, r, s)
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

func DbConnectionHealthHandler(w http.ResponseWriter, r *http.Request, s services.Service) {
	jsonResp, err := json.Marshal(s.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
