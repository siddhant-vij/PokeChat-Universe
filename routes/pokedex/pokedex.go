package pokedexroutes

import (
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
)

func ServePokedexPage(w http.ResponseWriter, r *http.Request) {
	pokedexPage := pages.PokedexPage()
	pokedexPage.Render(r.Context(), w)
}
