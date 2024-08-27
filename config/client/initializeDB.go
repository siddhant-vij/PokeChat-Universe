package client

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/database"
)

type pokemonFromAPI struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Height  int    `json:"height"`
	Weight  int    `json:"weight"`
	Sprites struct {
		Other struct {
			OfficialArtwork struct {
				PictureUrl string `json:"front_default"`
			} `json:"official-artwork"`
		} `json:"other"`
	} `json:"sprites"`
	BaseExperience int `json:"base_experience"`
	Types          []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
	} `json:"stats"`
	Species struct {
		Url string `json:"url"`
	} `json:"species"`
}

type speciesDataFromAPI struct {
	EvolutionChain struct {
		Url string `json:"url"`
	} `json:"evolution_chain"`
}

type evolutionChainFromAPI struct {
	Chain evolutionNode `json:"chain"`
}

type evolutionNode struct {
	Species struct {
		URL string `json:"url"`
	} `json:"species"`
	EvolvesTo []evolutionNode `json:"evolves_to"`
}

func FetchAndInsertRequest(cfg *config.AppConfig) {
	if toFetchFromAPI(cfg) {
		var wg sync.WaitGroup

		for _, id := range pokemonIdsToFetch {
			wg.Add(1)
			go func(pokemonID int) {
				defer wg.Done()
				err := fetchDataAndInsertIntoDB(cfg, pokemonID)
				if err != nil {
					log.Printf("error fetching and inserting pokemon with id: %d into DB. Err: %v", pokemonID, err)
					return
				}
			}(id)
		}

		wg.Wait()
		log.Println("Database Initialized!")
	} else {
		log.Println("Database initialized. No new pokemon to fetch!")
	}
	pokemonIdsToFetch = []int{}
}

func fetchDataAndInsertIntoDB(cfg *config.AppConfig, pokemonID int) error {
	// Stage 1: Fetch & Insert Pokemon Data
	var pokemonData pokemonFromAPI
	err := do(fmt.Sprintf("/pokemon/%d", pokemonID), &pokemonData)
	if err != nil {
		return err
	}
	err = insertPokemonDataIntoDB(cfg, pokemonData)
	if err != nil {
		return err
	}

	// Stage 2: Fetch Chain ID from Species URL
	speciesId := getIdFromUrl(pokemonData.Species.Url)
	var speciesData speciesDataFromAPI
	err = do(fmt.Sprintf("/pokemon-species/%d", speciesId), &speciesData)
	if err != nil {
		return err
	}
	chainId := getIdFromUrl(speciesData.EvolutionChain.Url)

	// Stage 3: Fetch & Insert Evolution Chain Data
	var evolutionChainData evolutionChainFromAPI
	err = do(fmt.Sprintf("/evolution-chain/%d", chainId), &evolutionChainData)
	if err != nil {
		return err
	}
	err = insertEvolutionChainDataIntoDB(cfg, evolutionChainData, chainId)
	if err != nil {
		return err
	}
	return nil
}

func insertPokemonDataIntoDB(cfg *config.AppConfig, pokemonData pokemonFromAPI) error {
	var types []string
	for _, t := range pokemonData.Types {
		types = append(types, t.Type.Name)
	}

	var insertPokemonParams = database.InsertPokemonParams{
		ID:             int32(pokemonData.Id),
		Name:           pokemonData.Name,
		Height:         int32(pokemonData.Height),
		Weight:         int32(pokemonData.Weight),
		PictureUrl:     pokemonData.Sprites.Other.OfficialArtwork.PictureUrl,
		BaseExperience: int32(pokemonData.BaseExperience),
		Types:          types,
		Hp:             int32(pokemonData.Stats[0].BaseStat),
		Attack:         int32(pokemonData.Stats[1].BaseStat),
		Defense:        int32(pokemonData.Stats[2].BaseStat),
		SpecialAttack:  int32(pokemonData.Stats[3].BaseStat),
		SpecialDefense: int32(pokemonData.Stats[4].BaseStat),
		Speed:          int32(pokemonData.Stats[5].BaseStat),
	}

	cfg.Mutex.Lock()
	err := cfg.DBQueries.InsertPokemon(context.Background(), insertPokemonParams)
	cfg.Mutex.Unlock()
	if err != nil {
		return err
	}
	return nil
}

func insertEvolutionChainDataIntoDB(cfg *config.AppConfig, evolutionChainData evolutionChainFromAPI, chainId int) error {
	var pokemonEvolutionIDs []int
	extractPokemonIDs(&evolutionChainData.Chain, &pokemonEvolutionIDs)

	for i := 0; i < len(pokemonEvolutionIDs); i++ {
		var insertPokemonEvolutionParams = database.InsertPokemonEvolutionParams{
			ChainID:   int32(chainId),
			PokemonID: int32(pokemonEvolutionIDs[i]),
		}

		if i+1 < len(pokemonEvolutionIDs) {
			insertPokemonEvolutionParams.EvolvesToID = sql.NullInt32{
				Int32: int32(pokemonEvolutionIDs[i+1]),
				Valid: true,
			}
		} else {
			insertPokemonEvolutionParams.EvolvesToID = sql.NullInt32{
				Valid: false,
			}
		}

		cfg.Mutex.Lock()
		err := cfg.DBQueries.InsertPokemonEvolution(context.Background(), insertPokemonEvolutionParams)
		cfg.Mutex.Unlock()
		if err != nil {
			return err
		}
	}
	return nil
}

func extractPokemonIDs(node *evolutionNode, pokemonIds *[]int) {
	for _, nextNode := range node.EvolvesTo {
		extractPokemonIDs(&nextNode, pokemonIds)
	}

	*pokemonIds = append(*pokemonIds, getIdFromUrl(node.Species.URL))
}
