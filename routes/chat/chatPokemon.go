package chatroutes

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"

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

	insertChatParams := pokedex.InsertChatEntryParams{
		ID:        uuid.New(),
		UserID:    cfg.LoggedInUserId,
		PokemonID: int32(pokemon.ID),
	}
	err = cfg.DBQueries.InsertChatEntry(context.Background(), insertChatParams)
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

func AddPokemonToPokedexFromChatWindow(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
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

	insertChatParams := pokedex.InsertChatEntryParams{
		ID:        uuid.New(),
		UserID:    cfg.LoggedInUserId,
		PokemonID: int32(pokemonId),
	}
	err = cfg.DBQueries.InsertChatEntry(context.Background(), insertChatParams)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	pokemon, err := cfg.DBQueries.GetPokemonByID(context.Background(), int32(pokemonId))
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}
	pokemonName := utils.FormatName(pokemon.Name)

	addChatUpdate := pages.ChatWindowUpdateOnAdd(pokemonName)
	addChatUpdate.Render(r.Context(), w)

	buttonUpdateOOB := pages.ChatPokemonFooterAddOOB(int(pokemon.ID))
	buttonUpdateOOB.Render(r.Context(), w)
}

func RemovePokemonFromPokedexInChatWindow(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	pokemonIdStr := r.FormValue("pokemonIdStr")
	pokemonId, err := utils.DeformatId(pokemonIdStr)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	deleteParams := pokedex.DeleteUserCollectedPokemonParams{
		UserID:    cfg.LoggedInUserId,
		PokemonID: int32(pokemonId),
	}
	err = cfg.DBQueries.DeleteUserCollectedPokemon(context.Background(), deleteParams)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	pokemon, err := cfg.DBQueries.GetPokemonByID(context.Background(), int32(pokemonId))
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}
	pokemonName := utils.FormatName(pokemon.Name)

	removeChatUpdate := pages.ChatWindowUpdateOnRemove(pokemonName)
	removeChatUpdate.Render(r.Context(), w)

	buttonUpdateOOB := pages.ChatPokemonFooterRemoveOOB(int(pokemon.ID))
	buttonUpdateOOB.Render(r.Context(), w)
}
