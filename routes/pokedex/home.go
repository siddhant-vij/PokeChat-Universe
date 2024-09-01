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

var currentOffset int

func ServeHomePage(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	currentOffset = 0
	params := pokedex.GetPokemonsSortedByIdAscParams{
		Limit:  12,
		Offset: int32(currentOffset),
	}
	pokemonList, err := cfg.DBQueries.GetPokemonsSortedByIdAsc(context.Background(), params)
	if err != nil {
		log.Fatalf("error getting the initial pokemon list from DB - Home Available Tab. Err: %v", err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	var homeAvailablePokemons = make([]utils.HomeAvailablePokemon, 0)

	for _, pokemon := range pokemonList {
		homeAvailablePokemons = append(homeAvailablePokemons, utils.HomeAvailablePokemon{
			ID:         utils.FormatID(int(pokemon.ID)),
			Name:       utils.FormatName(pokemon.Name),
			PictureUrl: pokemon.PictureUrl,
			Types:      utils.DisplayTypes(pokemon.Types),
		})
	}

	currentOffset += 12

	homePage := pages.HomePage(homeAvailablePokemons)
	homePage.Render(r.Context(), w)
}

func HomeAvailableLoadMore(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	params := pokedex.GetPokemonsSortedByIdAscParams{
		Limit:  12,
		Offset: int32(currentOffset),
	}
	pokemonList, err := cfg.DBQueries.GetPokemonsSortedByIdAsc(context.Background(), params)
	if err != nil {
		log.Fatalf("error getting the load more pokemon list from DB - Home Available Tab at offset: %d. Err: %v", currentOffset, err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	var homeAvailablePokemons = make([]utils.HomeAvailablePokemon, 0)

	for _, pokemon := range pokemonList {
		homeAvailablePokemons = append(homeAvailablePokemons, utils.HomeAvailablePokemon{
			ID:         utils.FormatID(int(pokemon.ID)),
			Name:       utils.FormatName(pokemon.Name),
			PictureUrl: pokemon.PictureUrl,
			Types:      utils.DisplayTypes(pokemon.Types),
		})
	}

	currentOffset += 12

	for _, pokemon := range homeAvailablePokemons {
		pokemonCard := pages.HomeAvailablePokemonCard(pokemon)
		pokemonCard.Render(r.Context(), w)
	}

	if len(pokemonList) < 12 {
		loadMoreBtnDisabled := pages.LoadMoreButtonDisabled()
		loadMoreBtnDisabled.Render(r.Context(), w)
	}
}
