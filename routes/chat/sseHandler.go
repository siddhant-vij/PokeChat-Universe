package chatroutes

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/chat"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex/utils"
)

var (
	stopChannel   = make(chan struct{})
	stopChannelMu sync.Mutex
)

func SseHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	clientGone := r.Context().Done()

	pokemonName := r.URL.Query().Get("pokemonName")

	userMessage := r.URL.Query().Get("userMessage")
	llmResponseStream := chat.GenerateChatResponse(userMessage, pokemonName, cfg)

	var pokemonResponse string

outerLoop:
	for {
		select {
		case <-clientGone:
			fmt.Println("Client disconnected")
			break outerLoop
		case <-stopChannel:
			fmt.Fprintf(w, "event: Complete\n")
			fmt.Fprintf(w, "data: Stream stopped\n\n")
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			break outerLoop
		case nextToken, ok := <-llmResponseStream:
			if !ok {
				fmt.Fprintf(w, "event: Complete\n")
				fmt.Fprintf(w, "data: LLM simulation done!\n\n")
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
				break outerLoop
			}

			pokemonResponse += nextToken

			fmt.Fprintf(w, "data: %s\n\n", nextToken)

			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		}
	}

	pokemon, err := cfg.DBQueries.GetPokemonDetailsByName(r.Context(), utils.DeformatName(pokemonName))
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	insertChatHistoryParams := pokedex.InsertChatHistoryParams{
		ID:        uuid.New(),
		UserID:    cfg.LoggedInUserId,
		PokemonID: int32(pokemon.ID),
		Sender:    "ai",
		Message:   pokemonResponse,
	}
	err = cfg.DBQueries.InsertChatHistory(r.Context(), insertChatHistoryParams)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}
}

func StopSseHandler(w http.ResponseWriter, r *http.Request) {
	stopChannelMu.Lock()
	defer stopChannelMu.Unlock()
	close(stopChannel)
	stopChannel = make(chan struct{})
}
