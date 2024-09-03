package pokedexroutes

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

func CollectPokemonHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
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

	var newPokemonDisplay utils.PokemonDisplay

	sortCriteria := r.FormValue("sortBy")
	switch sortCriteria {
	case "id-asc":
		lastFetchedPokemonID, _ := utils.DeformatId(lastFetchedPokemon["id-asc"])

		newPokemonParams := pokedex.GetOneAvailablePokemonAfterCollectionByIdAscParams{
			UserID: cfg.LoggedInUserId,
			ID:     int32(lastFetchedPokemonID),
		}
		newPokemon, err := cfg.DBQueries.GetOneAvailablePokemonAfterCollectionByIdAsc(context.Background(), newPokemonParams)
		if err != nil {
			log.Println(err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}

		lastFetchedPokemon["id-asc"] = utils.FormatID(int(newPokemon.ID))

		newPokemonDisplay = utils.PokemonDisplay{
			ID:         utils.FormatID(int(newPokemon.ID)),
			Name:       utils.FormatName(newPokemon.Name),
			PictureUrl: newPokemon.PictureUrl,
			Types:      utils.DisplayTypes(newPokemon.Types),
		}

	case "id-desc":
		lastFetchedPokemonID, _ := utils.DeformatId(lastFetchedPokemon["id-desc"])

		newPokemonParams := pokedex.GetOneAvailablePokemonAfterCollectionByIdDescParams{
			UserID: cfg.LoggedInUserId,
			ID:     int32(lastFetchedPokemonID),
		}
		newPokemon, err := cfg.DBQueries.GetOneAvailablePokemonAfterCollectionByIdDesc(context.Background(), newPokemonParams)
		if err != nil {
			log.Println(err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}

		lastFetchedPokemon["id-desc"] = utils.FormatID(int(newPokemon.ID))

		newPokemonDisplay = utils.PokemonDisplay{
			ID:         utils.FormatID(int(newPokemon.ID)),
			Name:       utils.FormatName(newPokemon.Name),
			PictureUrl: newPokemon.PictureUrl,
			Types:      utils.DisplayTypes(newPokemon.Types),
		}

	case "name-asc":
		lastFetchedPokemonName := utils.DeformatName(lastFetchedPokemon["name-asc"])

		newPokemonParams := pokedex.GetOneAvailablePokemonAfterCollectionByNameAscParams{
			UserID: cfg.LoggedInUserId,
			Name:   lastFetchedPokemonName,
		}
		newPokemon, err := cfg.DBQueries.GetOneAvailablePokemonAfterCollectionByNameAsc(context.Background(), newPokemonParams)
		if err != nil {
			log.Println(err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}

		lastFetchedPokemon["name-asc"] = utils.FormatName(newPokemon.Name)

		newPokemonDisplay = utils.PokemonDisplay{
			ID:         utils.FormatID(int(newPokemon.ID)),
			Name:       utils.FormatName(newPokemon.Name),
			PictureUrl: newPokemon.PictureUrl,
			Types:      utils.DisplayTypes(newPokemon.Types),
		}

	case "name-desc":
		lastFetchedPokemonName := utils.DeformatName(lastFetchedPokemon["name-desc"])

		newPokemonParams := pokedex.GetOneAvailablePokemonAfterCollectionByNameDescParams{
			UserID: cfg.LoggedInUserId,
			Name:   lastFetchedPokemonName,
		}
		newPokemon, err := cfg.DBQueries.GetOneAvailablePokemonAfterCollectionByNameDesc(context.Background(), newPokemonParams)
		if err != nil {
			log.Println(err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}

		lastFetchedPokemon["name-desc"] = utils.FormatName(newPokemon.Name)

		newPokemonDisplay = utils.PokemonDisplay{
			ID:         utils.FormatID(int(newPokemon.ID)),
			Name:       utils.FormatName(newPokemon.Name),
			PictureUrl: newPokemon.PictureUrl,
			Types:      utils.DisplayTypes(newPokemon.Types),
		}
	}

	newPokemonCard := pages.PokedexAvailableAfterCollection(newPokemonDisplay, sortCriteria)
	newPokemonCard.Render(r.Context(), w)
}
