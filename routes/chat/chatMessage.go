package chatroutes

import (
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
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

func ChatMessageButtonHandler(w http.ResponseWriter, r *http.Request) {
	userMessage := r.FormValue("userMessage")
	userMsgChat := pages.UserMessage(userMessage)
	userMsgChat.Render(r.Context(), w)

	pokemonName := r.FormValue("pokemonName")
	emptyChatOOB := pages.EmptyInputFormPostSendOOB(pokemonName)
	emptyChatOOB.Render(r.Context(), w)

	responseBtn := pages.ResponseSendButtonOOB()
	responseBtn.Render(r.Context(), w)

	// Response Generation via SSE...
}

func ResetButtonHandler(w http.ResponseWriter, r *http.Request) {
	disabledBtn := pages.DisabledSendButtonOOB()
	disabledBtn.Render(r.Context(), w)
}
