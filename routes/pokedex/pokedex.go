package pokedexroutes

import (
	"context"
	"log"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex/utils"
)

var (
	currentAvailableOffset int
	lastFetchedPokemon     map[string]string
)

func ServePokedexPage(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	isHtmxRequest := r.Header.Get("HX-Request") == "true"
	if isHtmxRequest {
		w.Header().Set("HX-Redirect", "/pokedex")
	} else {
		ServeAvailablePage(w, r, cfg)
	}
}

func ServeAvailablePage(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	currentAvailableOffset = 0
	lastFetchedPokemon = make(map[string]string)
	initialLimit := 12

	params := pokedex.GetUserAvailablePokemonsSortedByIdAscParams{
		UserID: cfg.LoggedInUserId,
		Limit:  int32(initialLimit),
		Offset: int32(currentAvailableOffset),
	}
	pokemonList, err := cfg.DBQueries.GetUserAvailablePokemonsSortedByIdAsc(context.Background(), params)
	if err != nil {
		log.Fatalf("error getting the initial user available pokemon list from DB - Serve Pokedex Available Tab. Err: %v", err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	var paPokemons = make([]utils.PokemonDisplay, 0)

	for id, pokemon := range pokemonList {
		if id == len(pokemonList)-1 {
			lastFetchedPokemon["id-asc"] = utils.FormatID(int(pokemon.ID))
		}
		paPokemons = append(paPokemons, utils.PokemonDisplay{
			ID:         utils.FormatID(int(pokemon.ID)),
			Name:       utils.FormatName(pokemon.Name),
			PictureUrl: pokemon.PictureUrl,
			Types:      utils.DisplayTypes(pokemon.Types),
		})
	}

	currentAvailableOffset += initialLimit

	pokedexPage := pages.PokedexPage(paPokemons)
	pokedexPage.Render(r.Context(), w)
}

func ServeCollectedPage(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	currentCollectedOffset = 0
	initialLimit := 12

	params := pokedex.GetUserCollectedPokemonsSortedByIdAscParams{
		UserID: cfg.LoggedInUserId,
		Limit:  int32(initialLimit),
		Offset: int32(currentCollectedOffset),
	}
	pokemonList, err := cfg.DBQueries.GetUserCollectedPokemonsSortedByIdAsc(context.Background(), params)
	if err != nil {
		log.Fatalf("error getting the initial user collected pokemon list from DB - Pokedex Collected Tab. Err: %v", err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	var pcPokemons = make([]utils.PokemonDisplay, 0)

	for _, pokemon := range pokemonList {
		pcPokemons = append(pcPokemons, utils.PokemonDisplay{
			ID:         utils.FormatID(int(pokemon.ID)),
			Name:       utils.FormatName(pokemon.Name),
			PictureUrl: pokemon.PictureUrl,
			Types:      utils.DisplayTypes(pokemon.Types),
		})
	}

	currentCollectedOffset += initialLimit

	pokedexCollectedPage := pages.PokedexCollectedPage(pcPokemons)
	pokedexCollectedPage.Render(r.Context(), w)
}
