package pokedexroutes

import (
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
)

func ServePokedexPage(w http.ResponseWriter, r *http.Request) {
	pokedexPage := pages.PokedexPage()
	pokedexPage.Render(r.Context(), w)
}

func AvailablePokedexHandler(w http.ResponseWriter, r *http.Request) {
	availableTab := pages.AvailableTemp()
	availableTab.Render(r.Context(), w)
}

func CollectedPokedexHandler(w http.ResponseWriter, r *http.Request) {
	collectedTab := pages.CollectedTemp()
	collectedTab.Render(r.Context(), w)
}

func ChatPokedexHandler(w http.ResponseWriter, r *http.Request) {
	chatTab := pages.ChatTemp()
	chatTab.Render(r.Context(), w)
}
