package health

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/services"
)

func ServiceConnectionHealthHandler(w http.ResponseWriter, r *http.Request, s services.Service) {
	jsonResp, err := json.Marshal(s.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
