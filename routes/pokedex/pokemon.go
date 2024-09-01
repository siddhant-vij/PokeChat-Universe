package pokedexroutes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
	"github.com/siddhant-vij/PokeChat-Universe/config"
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

	pokemonId := pokemon.ID

	evolutionChain, err := cfg.DBQueries.GetFullEvolutionChain(context.Background(), pokemonId)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	if cfg.AuthStatus {
		// check if the user has already collected the pokemon - isCollected? & then proceed here...
	} else {
		pokemonPage := pages.PokemonPage(pokemon, false, false, evolutionChain)
		pokemonPage.Render(r.Context(), w)
	}
}
