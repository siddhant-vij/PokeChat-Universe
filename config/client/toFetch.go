package client

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/siddhant-vij/PokeChat-Universe/config"
)

type pokemonsFetch struct {
	Count   int            `json:"count"`
	Results []pokemonCount `json:"results"`
}

type pokemonCount struct {
	Url string `json:"url"`
}

var pokemonIdsToFetch []int

func getPokemonCountFromDB(cfg *config.AppConfig) int {
	cfg.Mutex.RLock()
	count, err := cfg.DBQueries.GetPokemonCount(context.Background())
	cfg.Mutex.RUnlock()
	if err != nil {
		log.Fatalf("error getting pokemon count from DB. Err: %v", err)
	}
	return int(count)
}

func getPokemonCountFromAPI() int {
	var pokemonsFetch pokemonsFetch
	err := do("/pokemon", &pokemonsFetch)
	if err != nil {
		log.Fatalf("error getting full count from API. Err: %v", err)
	}

	fullCount := pokemonsFetch.Count

	err = do(fmt.Sprintf("/pokemon?limit=%d", fullCount), &pokemonsFetch)
	if err != nil {
		log.Fatalf("error fetching pokemons from API. Err: %v", err)
	}

	for _, pokemon := range pokemonsFetch.Results {
		id := getIdFromUrl(pokemon.Url)
		if id < 10001 {
			// From PokeAPI Exploration...
			pokemonIdsToFetch = append(pokemonIdsToFetch, id)
		}
	}
	return len(pokemonIdsToFetch)
}

func getIdFromUrl(url string) int {
	url = url[:len(url)-1]
	idStr := url[strings.LastIndex(url, "/")+1:]
	pokemonID, _ := strconv.Atoi(idStr)
	return pokemonID
}

func toFetchFromAPI(cfg *config.AppConfig) bool {
	return getPokemonCountFromDB(cfg) != getPokemonCountFromAPI()
}
