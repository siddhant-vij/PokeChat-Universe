package chatroutes

import (
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
	"github.com/siddhant-vij/PokeChat-Universe/config"
)

func ChatWithPokemonHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	pokemonName := r.FormValue("pokemonName")
	chatPage := pages.PokedexChat(pokemonName)
	chatPage.Render(r.Context(), w)

	chatHeader := pages.ChatHeaderUpdateOOB(cfg.AuthStatus)
	chatHeader.Render(r.Context(), w)
}
