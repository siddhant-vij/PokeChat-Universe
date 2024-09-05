package pokedexroutes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex/utils"
)

var currentCollectedOffset int

func CollectedPokedexHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
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

	pokedexCollected := pages.PokedexCollected(pcPokemons)
	pokedexCollected.Render(r.Context(), w)
}

func PokedexCollectedSort(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	sortCriteria := r.FormValue("sort-by")
	var pcPokemons = make([]utils.PokemonDisplay, 0)
	initialLimit := 12

	switch sortCriteria {
	case "id-asc":
		currentCollectedOffset = 0
		params := pokedex.GetUserCollectedPokemonsSortedByIdAscParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(initialLimit),
			Offset: int32(currentCollectedOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserCollectedPokemonsSortedByIdAsc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the initial user collected pokemon list from DB - Pokedex Collected Tab: ID Asc. Err: %v", err)
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
		loadMoreBtn := pages.PCLoadMoreIdAscButton()
		loadMoreBtn.Render(r.Context(), w)
	case "id-desc":
		currentCollectedOffset = 0
		params := pokedex.GetUserCollectedPokemonsSortedByIdDescParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(initialLimit),
			Offset: int32(currentCollectedOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserCollectedPokemonsSortedByIdDesc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the initial user collected pokemon list from DB - Pokedex Collected Tab: ID Desc. Err: %v", err)
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
		loadMoreBtn := pages.PCLoadMoreIdDescButton()
		loadMoreBtn.Render(r.Context(), w)
	case "name-asc":
		currentCollectedOffset = 0
		params := pokedex.GetUserCollectedPokemonsSortedByNameAscParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(initialLimit),
			Offset: int32(currentCollectedOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserCollectedPokemonsSortedByNameAsc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the initial user collected pokemon list from DB - Pokedex Collected Tab: Name Asc. Err: %v", err)
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
		loadMoreBtn := pages.PCLoadMoreNameAscButton()
		loadMoreBtn.Render(r.Context(), w)
	case "name-desc":
		currentCollectedOffset = 0
		params := pokedex.GetUserCollectedPokemonsSortedByNameDescParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(initialLimit),
			Offset: int32(currentCollectedOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserCollectedPokemonsSortedByNameDesc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the initial user collected pokemon list from DB - Pokedex Collected Tab: Name Desc. Err: %v", err)
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
		loadMoreBtn := pages.PCLoadMoreNameDescButton()
		loadMoreBtn.Render(r.Context(), w)
	}

	currentCollectedOffset += initialLimit

	for _, pokemon := range pcPokemons {
		pokemonCard := pages.PokedexCollectedPokemonCard(pokemon, sortCriteria)
		pokemonCard.Render(r.Context(), w)
	}

	searchUpdate := pages.PCSearchUpdateOOB()
	searchUpdate.Render(r.Context(), w)

	if len(pcPokemons) < initialLimit {
		loadMoreBtnDisabled := pages.PCLoadMoreButtonDisabled()
		loadMoreBtnDisabled.Render(r.Context(), w)
	}
}

func PokedexCollectedLoadMore(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	sortCriteria := r.FormValue("sort-by")
	loadMoreLimit := 12
	var pcPokemons = make([]utils.PokemonDisplay, 0)

	switch sortCriteria {
	case "id-asc":
		params := pokedex.GetUserCollectedPokemonsSortedByIdAscParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(loadMoreLimit),
			Offset: int32(currentCollectedOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserCollectedPokemonsSortedByIdAsc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the load more pokemon list from DB - Pokedex Collected Tab at offset: %d - ID Asc. Err: %v", currentCollectedOffset, err)
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
	case "id-desc":
		params := pokedex.GetUserCollectedPokemonsSortedByIdDescParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(loadMoreLimit),
			Offset: int32(currentCollectedOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserCollectedPokemonsSortedByIdDesc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the load more pokemon list from DB - Pokedex Collected Tab at offset: %d - ID Desc. Err: %v", currentCollectedOffset, err)
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
	case "name-asc":
		params := pokedex.GetUserCollectedPokemonsSortedByNameAscParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(loadMoreLimit),
			Offset: int32(currentCollectedOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserCollectedPokemonsSortedByNameAsc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the load more pokemon list from DB - Pokedex Collected Tab at offset: %d - Name Asc. Err: %v", currentCollectedOffset, err)
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
	case "name-desc":
		params := pokedex.GetUserCollectedPokemonsSortedByNameDescParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(loadMoreLimit),
			Offset: int32(currentCollectedOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserCollectedPokemonsSortedByNameDesc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the load more pokemon list from DB - Pokedex Collected Tab at offset: %d - Name Desc. Err: %v", currentCollectedOffset, err)
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
	}

	currentCollectedOffset += loadMoreLimit

	for _, pokemon := range pcPokemons {
		pokemonCard := pages.PokedexCollectedPokemonCard(pokemon, sortCriteria)
		pokemonCard.Render(r.Context(), w)
	}

	if len(pcPokemons) < loadMoreLimit {
		loadMoreBtnDisabled := pages.PCLoadMoreButtonDisabled()
		loadMoreBtnDisabled.Render(r.Context(), w)
	}
}

func PokedexCollectedSearch(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	pokemonName := r.FormValue("pokemonName")
	if pokemonName == "" {
		r.Form.Add("sort-by", "id-asc")
		PokedexCollectedSort(w, r, cfg)
		sortUpdate := pages.PCSortUpdateSelectedOOB()
		sortUpdate.Render(r.Context(), w)
		return
	}

	params := pokedex.SearchUserCollectedPokemonsByNameParams{
		UserID: cfg.LoggedInUserId,
		Name:   fmt.Sprintf("%s%%", pokemonName),
		Limit:  12,
	}

	pokemonList, err := cfg.DBQueries.SearchUserCollectedPokemonsByName(context.Background(), params)
	if err != nil {
		log.Fatalf("error getting the search pokemon list from DB - Pokedex Collected Tab. Err: %v", err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	var pcSearchPokemons = make([]utils.PokemonDisplay, 0)

	for id, pokemon := range pokemonList {
		if id == len(pokemonList)-1 {
			lastFetchedPokemon["name-asc"] = utils.FormatName(pokemon.Name)
		}
		pcSearchPokemons = append(pcSearchPokemons, utils.PokemonDisplay{
			ID:         utils.FormatID(int(pokemon.ID)),
			Name:       utils.FormatName(pokemon.Name),
			PictureUrl: pokemon.PictureUrl,
			Types:      utils.DisplayTypes(pokemon.Types),
		})
	}

	for _, pokemon := range pcSearchPokemons {
		pokemonCard := pages.PokedexCollectedPokemonCard(pokemon, "name-asc")
		pokemonCard.Render(r.Context(), w)
	}

	sortUpdate := pages.PCSortUpdateOOB()
	sortUpdate.Render(r.Context(), w)

	loadMoreSearchBtnDisabled := pages.PCLoadMoreSearchButtonDisabled()
	loadMoreSearchBtnDisabled.Render(r.Context(), w)
}
