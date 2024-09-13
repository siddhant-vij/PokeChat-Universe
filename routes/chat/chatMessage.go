package chatroutes

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	
	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex/utils"
)

func ChatMessageHandler(w http.ResponseWriter, r *http.Request) {
	userMessage := r.FormValue("userMessage")
	if userMessage != "" {
		activeBtn := pages.ActiveSendButton()
		activeBtn.Render(r.Context(), w)
	} else {
		disabledBtn := pages.DisabledSendButton()
		disabledBtn.Render(r.Context(), w)
	}
}

func ChatMessageButtonHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	uniqueID := time.Now().UnixNano()

	userMessage := r.FormValue("userMessage")
	userMsgChat := pages.UserAndPokemonMessage(userMessage, uniqueID)
	userMsgChat.Render(r.Context(), w)

	pokemonName := r.FormValue("pokemonName")
	emptyChatOOB := pages.EmptyInputFormPostSendOOB(pokemonName, true)
	emptyChatOOB.Render(r.Context(), w)

	responseBtn := pages.ResponseSendButtonOOB(uniqueID, pokemonName, userMessage)
	responseBtn.Render(r.Context(), w)

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
		Sender:    "user",
		Message:   userMessage,
	}
	err = cfg.DBQueries.InsertChatHistory(r.Context(), insertChatHistoryParams)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}
}

func RenderButtonUpdate(w http.ResponseWriter, r *http.Request) {
	disabledBtn := pages.DisabledSendButton()
	disabledBtn.Render(r.Context(), w)

	pokemonName := r.FormValue("pokemonName")
	emptyChatOOB := pages.EmptyInputFormPostSendOOB(pokemonName, false)
	emptyChatOOB.Render(r.Context(), w)
}
