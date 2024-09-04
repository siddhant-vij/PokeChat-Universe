package pokedexroutes

import (
	"fmt"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
)

func ChatToAvailableRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Redirect", "/pokedex")
}

func AvailableToChatRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Redirect", "/pokeChat/")
}

func ChatToCollectedRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Redirect", "/collectedPokedex")
}

func CollectedToChatRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Redirect", "/pokeChat/")
}

func ChatPokedexHandler(w http.ResponseWriter, r *http.Request) {
	pokemonName := r.FormValue("pokemonName")
	if pokemonName == "" || pokemonName == "*" {
		chatPage := pages.PokedexChat("")
		chatPage.Render(r.Context(), w)
		return
	}
	w.Header().Set("HX-Redirect", fmt.Sprintf("/pokeChat/%s", pokemonName))
}

func ChatWithPokemonHandler(w http.ResponseWriter, r *http.Request) {
	pokemonName := r.URL.Path[10:]
	chatPage := pages.PokedexChatPage(pokemonName)
	chatPage.Render(r.Context(), w)
}
