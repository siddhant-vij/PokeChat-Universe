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

var currentOffset int

func ServeHomePage(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	currentOffset = 0
	initialLimit := 12

	params := pokedex.GetPokemonsSortedByIdAscParams{
		Limit:  int32(initialLimit),
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

	currentOffset += initialLimit

	homePage := pages.HomePage(homeAvailablePokemons)
	homePage.Render(r.Context(), w)
}

func HomeAvailableSort(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	sortCriteria := r.FormValue("sort-by")
	var homeAvailablePokemons = make([]utils.HomeAvailablePokemon, 0)
	initialLimit := 12

	switch sortCriteria {
	case "id-asc":
		currentOffset = 0
		params := pokedex.GetPokemonsSortedByIdAscParams{
			Limit:  int32(initialLimit),
			Offset: int32(currentOffset),
		}
		pokemonList, err := cfg.DBQueries.GetPokemonsSortedByIdAsc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the initial pokemon list from DB - Home Available Tab - ID Asc. Err: %v", err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for _, pokemon := range pokemonList {
			homeAvailablePokemons = append(homeAvailablePokemons, utils.HomeAvailablePokemon{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
		loadMoreBtn := pages.LoadMoreIdAscButton()
		loadMoreBtn.Render(r.Context(), w)
	case "id-desc":
		currentOffset = 0
		params := pokedex.GetPokemonsSortedByIdDescParams{
			Limit:  int32(initialLimit),
			Offset: int32(currentOffset),
		}
		pokemonList, err := cfg.DBQueries.GetPokemonsSortedByIdDesc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the initial pokemon list from DB - Home Available Tab - ID Desc. Err: %v", err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for _, pokemon := range pokemonList {
			homeAvailablePokemons = append(homeAvailablePokemons, utils.HomeAvailablePokemon{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
		loadMoreBtn := pages.LoadMoreIdDescButton()
		loadMoreBtn.Render(r.Context(), w)
	case "name-asc":
		currentOffset = 0
		params := pokedex.GetPokemonsSortedByNameAscParams{
			Limit:  int32(initialLimit),
			Offset: int32(currentOffset),
		}
		pokemonList, err := cfg.DBQueries.GetPokemonsSortedByNameAsc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the initial pokemon list from DB - Home Available Tab - Name Asc. Err: %v", err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for _, pokemon := range pokemonList {
			homeAvailablePokemons = append(homeAvailablePokemons, utils.HomeAvailablePokemon{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
		loadMoreBtn := pages.LoadMoreNameAscButton()
		loadMoreBtn.Render(r.Context(), w)
	case "name-desc":
		currentOffset = 0
		params := pokedex.GetPokemonsSortedByNameDescParams{
			Limit:  int32(initialLimit),
			Offset: int32(currentOffset),
		}
		pokemonList, err := cfg.DBQueries.GetPokemonsSortedByNameDesc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the initial pokemon list from DB - Home Available Tab - Name Desc. Err: %v", err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for _, pokemon := range pokemonList {
			homeAvailablePokemons = append(homeAvailablePokemons, utils.HomeAvailablePokemon{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
		loadMoreBtn := pages.LoadMoreNameDescButton()
		loadMoreBtn.Render(r.Context(), w)
	}

	currentOffset += initialLimit

	for _, pokemon := range homeAvailablePokemons {
		pokemonCard := pages.HomeAvailablePokemonCard(pokemon)
		pokemonCard.Render(r.Context(), w)
	}
}

func HomeAvailableLoadMore(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	sortCriteria := r.FormValue("sort-by")
	loadMoreLimit := 12
	var homeAvailablePokemons = make([]utils.HomeAvailablePokemon, 0)

	switch sortCriteria {
	case "id-asc":
		params := pokedex.GetPokemonsSortedByIdAscParams{
			Limit:  int32(loadMoreLimit),
			Offset: int32(currentOffset),
		}
		pokemonList, err := cfg.DBQueries.GetPokemonsSortedByIdAsc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the load more pokemon list from DB - Home Available Tab at offset: %d - ID Asc. Err: %v", currentOffset, err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for _, pokemon := range pokemonList {
			homeAvailablePokemons = append(homeAvailablePokemons, utils.HomeAvailablePokemon{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
	case "id-desc":
		params := pokedex.GetPokemonsSortedByIdDescParams{
			Limit:  int32(loadMoreLimit),
			Offset: int32(currentOffset),
		}
		pokemonList, err := cfg.DBQueries.GetPokemonsSortedByIdDesc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the load more pokemon list from DB - Home Available Tab at offset: %d - ID Desc. Err: %v", currentOffset, err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for _, pokemon := range pokemonList {
			homeAvailablePokemons = append(homeAvailablePokemons, utils.HomeAvailablePokemon{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
	case "name-asc":
		params := pokedex.GetPokemonsSortedByNameAscParams{
			Limit:  int32(loadMoreLimit),
			Offset: int32(currentOffset),
		}
		pokemonList, err := cfg.DBQueries.GetPokemonsSortedByNameAsc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the load more pokemon list from DB - Home Available Tab at offset: %d - Name Asc. Err: %v", currentOffset, err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for _, pokemon := range pokemonList {
			homeAvailablePokemons = append(homeAvailablePokemons, utils.HomeAvailablePokemon{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
	case "name-desc":
		params := pokedex.GetPokemonsSortedByNameDescParams{
			Limit:  int32(loadMoreLimit),
			Offset: int32(currentOffset),
		}
		pokemonList, err := cfg.DBQueries.GetPokemonsSortedByNameDesc(context.Background(), params)
		if err != nil {
			log.Fatalf("error getting the load more pokemon list from DB - Home Available Tab at offset: %d - Name Desc. Err: %v", currentOffset, err)
			serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
			serverErrorPage.Render(r.Context(), w)
			return
		}
		for _, pokemon := range pokemonList {
			homeAvailablePokemons = append(homeAvailablePokemons, utils.HomeAvailablePokemon{
				ID:         utils.FormatID(int(pokemon.ID)),
				Name:       utils.FormatName(pokemon.Name),
				PictureUrl: pokemon.PictureUrl,
				Types:      utils.DisplayTypes(pokemon.Types),
			})
		}
	}

	currentOffset += loadMoreLimit

	for _, pokemon := range homeAvailablePokemons {
		pokemonCard := pages.HomeAvailablePokemonCard(pokemon)
		pokemonCard.Render(r.Context(), w)
	}

	if len(homeAvailablePokemons) < loadMoreLimit {
		loadMoreBtnDisabled := pages.LoadMoreButtonDisabled()
		loadMoreBtnDisabled.Render(r.Context(), w)
	}
}

func HomeAvailableSearch(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	pokemonName := r.FormValue("pokemonName")
	if pokemonName == "" {
		w.Header().Set("HX-Redirect", "/pokedex")
		return
	}

	params := pokedex.SearchPokemonByNameParams{
		Name:  fmt.Sprintf("%s%%", pokemonName),
		Limit: 12,
	}

	pokemonList, err := cfg.DBQueries.SearchPokemonByName(context.Background(), params)
	if err != nil {
		log.Fatalf("error getting the search pokemon list from DB - Home Available Tab. Err: %v", err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	var homeSearchPokemons = make([]utils.HomeAvailablePokemon, 0)

	for _, pokemon := range pokemonList {
		homeSearchPokemons = append(homeSearchPokemons, utils.HomeAvailablePokemon{
			ID:         utils.FormatID(int(pokemon.ID)),
			Name:       utils.FormatName(pokemon.Name),
			PictureUrl: pokemon.PictureUrl,
			Types:      utils.DisplayTypes(pokemon.Types),
		})
	}

	for _, pokemon := range homeSearchPokemons {
		pokemonCard := pages.HomeAvailablePokemonCard(pokemon)
		pokemonCard.Render(r.Context(), w)
	}

	loadMoreSearchBtnDisabled := pages.LoadMoreSearchButtonDisabled()
	loadMoreSearchBtnDisabled.Render(r.Context(), w)
}
