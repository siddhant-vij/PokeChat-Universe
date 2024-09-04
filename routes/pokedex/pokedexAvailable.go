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

func AvailablePokedexHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	currentAvailableOffset = 0
	initialLimit := 12

	params := pokedex.GetUserAvailablePokemonsSortedByIdAscParams{
		UserID: cfg.LoggedInUserId,
		Limit:  int32(initialLimit),
		Offset: int32(currentAvailableOffset),
	}
	pokemonList, err := cfg.DBQueries.GetUserAvailablePokemonsSortedByIdAsc(context.Background(), params)
	if err != nil {
		log.Fatalf("error getting the initial user available pokemon list from DB - Pokedex Available Tab. Err: %v", err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	var paPokemons = make([]utils.PokemonDisplay, 0)

	for _, pokemon := range pokemonList {
		paPokemons = append(paPokemons, utils.PokemonDisplay{
			ID:         utils.FormatID(int(pokemon.ID)),
			Name:       utils.FormatName(pokemon.Name),
			PictureUrl: pokemon.PictureUrl,
			Types:      utils.DisplayTypes(pokemon.Types),
		})
	}

	currentAvailableOffset += initialLimit

	pokedexAvailable := pages.PokedexAvailable(paPokemons)
	pokedexAvailable.Render(r.Context(), w)
}
