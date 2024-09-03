package pokedexroutes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex"
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

	pokemonId := pokemon.ID

	evolutionChain, err := cfg.DBQueries.GetFullEvolutionChain(context.Background(), pokemonId)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	if cfg.AuthStatus {
		isCollectedParams := pokedex.IsPokemonCollectedParams{
			UserID:    cfg.LoggedInUserId,
			PokemonID: int32(pokemonId),
		}
		isCollected, err := cfg.DBQueries.IsPokemonCollected(context.Background(), isCollectedParams)
		if err != nil {
			log.Println(err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		pokemonPage := pages.PokemonPage(pokemon, cfg.AuthStatus, isCollected, evolutionChain)
		pokemonPage.Render(r.Context(), w)
	} else {
		pokemonPage := pages.PokemonPage(pokemon, cfg.AuthStatus, false, evolutionChain)
		pokemonPage.Render(r.Context(), w)
	}
}

func CollectPokemonHandlerOnPokemonPage(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	pokemonIdStr := r.FormValue("pokemonIdStr")
	pokemonId, err := utils.DeformatId(pokemonIdStr)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	insertParams := pokedex.InsertUserCollectedPokemonParams{
		ID:        uuid.New(),
		UserID:    cfg.LoggedInUserId,
		PokemonID: int32(pokemonId),
	}
	err = cfg.DBQueries.InsertUserCollectedPokemon(context.Background(), insertParams)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	pokemonName := r.FormValue("pokemonName")
	pokemonCollectedButton := pages.PokemonCollectedButton(pokemonName)
	pokemonCollectedButton.Render(r.Context(), w)
}
