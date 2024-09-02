package pokedexroutes

import (
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
)

func ChatPokedexHandler(w http.ResponseWriter, r *http.Request) {
	chatTab := pages.PokedexChat()
	chatTab.Render(r.Context(), w)
}
