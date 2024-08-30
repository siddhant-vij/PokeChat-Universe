package pokedexroutes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex/utils"
)

func GetPokemonHandler(w http.ResponseWriter, r *http.Request) {
	pokemonName := r.FormValue("pokemonName")
	w.Header().Set("HX-Redirect", fmt.Sprintf("/%s", pokemonName))
}

func ServePokemonPage(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	pokemonName := r.URL.Path[1:]
	pokemon, err := cfg.DBQueries.GetPokemonDetailsByName(context.Background(), pokemonName)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	pokemonNameDisplayTemp := utils.FormatName(pokemon.Name)
	pokemonPage := pages.PokemonPage(pokemonNameDisplayTemp, cfg.AuthStatus)
	pokemonPage.Render(r.Context(), w)
}
