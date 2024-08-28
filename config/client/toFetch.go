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

func getPokemonCountFromDB(cfg *config.AppConfig) (int, error) {
	count, err := cfg.DBQueries.GetPokemonCount(context.Background())
	if err != nil {
		return 0, fmt.Errorf("error getting pokemon count from DB: %w", err)
	}
	return int(count), nil
}

func getPokemonCountFromAPI() (int, error) {
	var pokemonsFetch pokemonsFetch
	err := do("/pokemon", &pokemonsFetch)
	if err != nil {
		return 0, fmt.Errorf("error getting full pokemon count from API: %w", err)
	}

	fullCount := pokemonsFetch.Count

	err = do(fmt.Sprintf("/pokemon?limit=%d", fullCount), &pokemonsFetch)
	if err != nil {
		return 0, fmt.Errorf("error getting pokemon results from API: %w", err)
	}

	for _, pokemon := range pokemonsFetch.Results {
		id := getIdFromUrl(pokemon.Url)
		if id < 10001 {
			// From PokeAPI Exploration...
			pokemonIdsToFetch = append(pokemonIdsToFetch, id)
		}
	}
	return len(pokemonIdsToFetch), nil
}

func getIdFromUrl(url string) int {
	url = url[:len(url)-1]
	idStr := url[strings.LastIndex(url, "/")+1:]
	pokemonID, _ := strconv.Atoi(idStr)
	return pokemonID
}

func toFetchFromAPI(cfg *config.AppConfig) bool {
	countFromDB, err := getPokemonCountFromDB(cfg)
	if err != nil {
		log.Printf("error checking if new pokemon needs to be fetched - DB Err: %v", err)
		return true
	}
	countFromAPI, err := getPokemonCountFromAPI()
	if err != nil {
		log.Printf("error checking if new pokemon needs to be fetched - API Err: %v", err)
		return true
	}
	return countFromDB != countFromAPI
}
