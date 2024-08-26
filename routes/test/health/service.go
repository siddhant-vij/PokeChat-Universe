package health

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/config"
)

func DatabaseConnectionHealthHandler(w http.ResponseWriter, r *http.Request, dbS *config.DbService) {
	jsonResp, err := json.Marshal(dbS.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal for database. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func RedisConnectionHealthHandler(w http.ResponseWriter, r *http.Request, rS *config.RedisService) {
	jsonResp, err := json.Marshal(rS.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal for redis. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
