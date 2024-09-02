package pokedexroutes

import (
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
)

func CollectedPokedexHandler(w http.ResponseWriter, r *http.Request) {
	collectedTab := pages.PokedexCollected()
	collectedTab.Render(r.Context(), w)
}
