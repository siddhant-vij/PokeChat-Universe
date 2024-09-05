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

func AvailablePokedexHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
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

	pokedexAvailable := pages.PokedexAvailable(paPokemons)
	pokedexAvailable.Render(r.Context(), w)
}

func PokedexAvailableSort(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	sortCriteria := r.FormValue("sort-by")
	var paPokemons = make([]utils.PokemonDisplay, 0)
	initialLimit := 12

	switch sortCriteria {
	case "id-asc":
		currentAvailableOffset = 0
		params := pokedex.GetUserAvailablePokemonsSortedByIdAscParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(initialLimit),
			Offset: int32(currentAvailableOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserAvailablePokemonsSortedByIdAsc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the initial user available pokemon list from DB - Pokedex Available Tab: ID Asc. Err: %v", err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
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
		loadMoreBtn := pages.PALoadMoreIdAscButton()
		loadMoreBtn.Render(r.Context(), w)
	case "id-desc":
		currentAvailableOffset = 0
		params := pokedex.GetUserAvailablePokemonsSortedByIdDescParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(initialLimit),
			Offset: int32(currentAvailableOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserAvailablePokemonsSortedByIdDesc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the initial user available pokemon list from DB - Pokedex Available Tab: ID Desc. Err: %v", err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for id, pokemon := range pokemonList {
			if id == len(pokemonList)-1 {
				lastFetchedPokemon["id-desc"] = utils.FormatID(int(pokemon.ID))
			}
			paPokemons = append(paPokemons, utils.PokemonDisplay{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
		loadMoreBtn := pages.PALoadMoreIdDescButton()
		loadMoreBtn.Render(r.Context(), w)
	case "name-asc":
		currentAvailableOffset = 0
		params := pokedex.GetUserAvailablePokemonsSortedByNameAscParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(initialLimit),
			Offset: int32(currentAvailableOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserAvailablePokemonsSortedByNameAsc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the initial user available pokemon list from DB - Pokedex Available Tab: Name Asc. Err: %v", err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for id, pokemon := range pokemonList {
			if id == len(pokemonList)-1 {
				lastFetchedPokemon["name-asc"] = utils.FormatName(pokemon.Name)
			}
			paPokemons = append(paPokemons, utils.PokemonDisplay{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
		loadMoreBtn := pages.PALoadMoreNameAscButton()
		loadMoreBtn.Render(r.Context(), w)
	case "name-desc":
		currentAvailableOffset = 0
		params := pokedex.GetUserAvailablePokemonsSortedByNameDescParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(initialLimit),
			Offset: int32(currentAvailableOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserAvailablePokemonsSortedByNameDesc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the initial user available pokemon list from DB - Pokedex Available Tab: Name Desc. Err: %v", err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for id, pokemon := range pokemonList {
			if id == len(pokemonList)-1 {
				lastFetchedPokemon["name-desc"] = utils.FormatName(pokemon.Name)
			}
			paPokemons = append(paPokemons, utils.PokemonDisplay{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
		loadMoreBtn := pages.PALoadMoreNameDescButton()
		loadMoreBtn.Render(r.Context(), w)
	}

	currentAvailableOffset += initialLimit

	for _, pokemon := range paPokemons {
		pokemonCard := pages.PokedexAvailablePokemonCard(pokemon, sortCriteria)
		pokemonCard.Render(r.Context(), w)
	}

	searchUpdate := pages.PASearchUpdateOOB()
	searchUpdate.Render(r.Context(), w)

	if len(paPokemons) < initialLimit {
		loadMoreBtnDisabled := pages.PALoadMoreButtonDisabled()
		loadMoreBtnDisabled.Render(r.Context(), w)
	}
}

func PokedexAvailableLoadMore(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	sortCriteria := r.FormValue("sort-by")
	loadMoreLimit := 12
	var paPokemons = make([]utils.PokemonDisplay, 0)

	switch sortCriteria {
	case "id-asc":
		params := pokedex.GetUserAvailablePokemonsSortedByIdAscParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(loadMoreLimit),
			Offset: int32(currentAvailableOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserAvailablePokemonsSortedByIdAsc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the load more pokemon list from DB - Pokedex Available Tab at offset: %d - ID Asc. Err: %v", currentAvailableOffset, err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
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
	case "id-desc":
		params := pokedex.GetUserAvailablePokemonsSortedByIdDescParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(loadMoreLimit),
			Offset: int32(currentAvailableOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserAvailablePokemonsSortedByIdDesc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the load more pokemon list from DB - Pokedex Available Tab at offset: %d - ID Desc. Err: %v", currentAvailableOffset, err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for id, pokemon := range pokemonList {
			if id == len(pokemonList)-1 {
				lastFetchedPokemon["id-desc"] = utils.FormatID(int(pokemon.ID))
			}
			paPokemons = append(paPokemons, utils.PokemonDisplay{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
	case "name-asc":
		params := pokedex.GetUserAvailablePokemonsSortedByNameAscParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(loadMoreLimit),
			Offset: int32(currentAvailableOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserAvailablePokemonsSortedByNameAsc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the load more pokemon list from DB - Pokedex Available Tab at offset: %d - Name Asc. Err: %v", currentAvailableOffset, err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for id, pokemon := range pokemonList {
			if id == len(pokemonList)-1 {
				lastFetchedPokemon["name-asc"] = utils.FormatName(pokemon.Name)
			}
			paPokemons = append(paPokemons, utils.PokemonDisplay{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
	case "name-desc":
		params := pokedex.GetUserAvailablePokemonsSortedByNameDescParams{
			UserID: cfg.LoggedInUserId,
			Limit:  int32(loadMoreLimit),
			Offset: int32(currentAvailableOffset),
		}
		pokemonList, err := cfg.DBQueries.GetUserAvailablePokemonsSortedByNameDesc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the load more pokemon list from DB - Pokedex Available Tab at offset: %d - Name Desc. Err: %v", currentAvailableOffset, err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for id, pokemon := range pokemonList {
			if id == len(pokemonList)-1 {
				lastFetchedPokemon["name-desc"] = utils.FormatName(pokemon.Name)
			}
			paPokemons = append(paPokemons, utils.PokemonDisplay{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
	}

	currentAvailableOffset += loadMoreLimit

	for _, pokemon := range paPokemons {
		pokemonCard := pages.PokedexAvailablePokemonCard(pokemon, sortCriteria)
		pokemonCard.Render(r.Context(), w)
	}

	if len(paPokemons) < loadMoreLimit {
		loadMoreBtnDisabled := pages.PALoadMoreButtonDisabled()
		loadMoreBtnDisabled.Render(r.Context(), w)
	}
}

func PokedexAvailableSearch(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	pokemonName := r.FormValue("pokemonName")
	if pokemonName == "" {
		r.Form.Add("sort-by", "id-asc")
		PokedexAvailableSort(w, r, cfg)
		sortUpdate := pages.PASortUpdateSelectedOOB()
		sortUpdate.Render(r.Context(), w)
		return
	}

	params := pokedex.SearchUserAvailablePokemonsByNameParams{
		UserID: cfg.LoggedInUserId,
		Name:   fmt.Sprintf("%s%%", pokemonName),
		Limit:  12,
	}

	pokemonList, err := cfg.DBQueries.SearchUserAvailablePokemonsByName(context.Background(), params)
	if err != nil {
		log.Fatalf("error getting the search pokemon list from DB - Pokedex Available Tab. Err: %v", err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	var paSearchPokemons = make([]utils.PokemonDisplay, 0)

	for id, pokemon := range pokemonList {
		if id == len(pokemonList)-1 {
			lastFetchedPokemon["name-asc"] = utils.FormatName(pokemon.Name)
		}
		paSearchPokemons = append(paSearchPokemons, utils.PokemonDisplay{
			ID:         utils.FormatID(int(pokemon.ID)),
			Name:       utils.FormatName(pokemon.Name),
			PictureUrl: pokemon.PictureUrl,
			Types:      utils.DisplayTypes(pokemon.Types),
		})
	}

	for _, pokemon := range paSearchPokemons {
		pokemonCard := pages.PokedexAvailablePokemonCard(pokemon, "name-asc")
		pokemonCard.Render(r.Context(), w)
	}

	sortUpdate := pages.PASortUpdateOOB()
	sortUpdate.Render(r.Context(), w)

	loadMoreSearchBtnDisabled := pages.PALoadMoreSearchButtonDisabled()
	loadMoreSearchBtnDisabled.Render(r.Context(), w)
}
