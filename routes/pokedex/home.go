package pokedexroutes

import (
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
)

func ServeHomePage(w http.ResponseWriter, r *http.Request) {
	homePage := pages.HomePage()
	homePage.Render(r.Context(), w)
}
