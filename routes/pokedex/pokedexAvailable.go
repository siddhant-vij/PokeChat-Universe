package pokedexroutes

import (
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
)

func AvailablePokedexHandler(w http.ResponseWriter, r *http.Request) {
	pokedexPage := pages.PokedexAvailable()
	pokedexPage.Render(r.Context(), w)
}
