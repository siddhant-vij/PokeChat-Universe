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

func ChatWithPokemonHandler(w http.ResponseWriter, r *http.Request, pokemonName string, cfg *config.AppConfig) {
	pokemon, err := cfg.DBQueries.GetPokemonDetailsByName(context.Background(), utils.DeformatName(pokemonName))
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	isCollectedParams := pokedex.IsPokemonCollectedParams{
		UserID:    cfg.LoggedInUserId,
		PokemonID: int32(pokemon.ID),
	}
	isCollected, err := cfg.DBQueries.IsPokemonCollected(context.Background(), isCollectedParams)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	chatPokemon := pages.PokedexChatPokemonPage(pokemon, isCollected)
	chatPokemon.Render(r.Context(), w)
}
