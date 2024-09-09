package chatroutes

import (
	"context"
	"log"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex/utils"
)

var currentChatCollectedOffset int

func PokedexChatHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	pokemonName := r.URL.Path[len("/chat/"):]
	if pokemonName == "" {
		ServeChatHomePage(w, r, cfg)
	} else {
		ChatWithPokemonHandler(w, r, pokemonName, cfg)
	}
}

func ServeChatHomePage(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	currentChatCollectedOffset = 0
	initialLimit := 12

	params := pokedex.GetUserCollectedPokemonsSortedByIdAscParams{
		UserID: cfg.LoggedInUserId,
		Limit:  int32(initialLimit),
		Offset: int32(currentChatCollectedOffset),
	}
	pokemonList, err := cfg.DBQueries.GetUserCollectedPokemonsSortedByIdAsc(context.Background(), params)
	if err != nil {
		log.Fatalf("error getting the initial user collected pokemon list from DB - Pokedex Chat Tab. Err: %v", err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	var pChatPokemons = make([]utils.PokemonDisplay, 0)

	for _, pokemon := range pokemonList {
		pChatPokemons = append(pChatPokemons, utils.PokemonDisplay{
			ID:         utils.FormatID(int(pokemon.ID)),
			Name:       utils.FormatName(pokemon.Name),
			PictureUrl: pokemon.PictureUrl,
			Types:      utils.DisplayTypes(pokemon.Types),
		})
	}

	currentChatCollectedOffset += initialLimit

	homeChat := pages.PokedexChatHome(pChatPokemons)
	homeChat.Render(r.Context(), w)
}

func PokedexChatLoadMoreHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	loadMoreLimit := 12
	var pcPokemons = make([]utils.PokemonDisplay, 0)

	params := pokedex.GetUserCollectedPokemonsSortedByIdAscParams{
		UserID: cfg.LoggedInUserId,
		Limit:  int32(loadMoreLimit),
		Offset: int32(currentChatCollectedOffset),
	}
	pokemonList, err := cfg.DBQueries.GetUserCollectedPokemonsSortedByIdAsc(context.Background(), params)
	if err != nil {
		log.Fatalf("error getting the load more pokemon list from DB - Pokedex Collected Tab at offset: %d - ID Asc. Err: %v", currentChatCollectedOffset, err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}
	for _, pokemon := range pokemonList {
		pcPokemons = append(pcPokemons, utils.PokemonDisplay{
			ID:         utils.FormatID(int(pokemon.ID)),
			Name:       utils.FormatName(pokemon.Name),
			PictureUrl: pokemon.PictureUrl,
			Types:      utils.DisplayTypes(pokemon.Types),
		})
	}

	currentChatCollectedOffset += loadMoreLimit

	for _, pokemon := range pcPokemons {
		pokemonCard := pages.ChatPokemonCard(pokemon)
		pokemonCard.Render(r.Context(), w)
	}

	if len(pcPokemons) < loadMoreLimit {
		loadMoreBtnDisabled := pages.ChatHomeLoadMoreBtnDisabled()
		loadMoreBtnDisabled.Render(r.Context(), w)
	}
}
