package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex"
)

func DeformatId(formattedId string) (int, error) {
	id, err := strconv.Atoi(strings.TrimPrefix(formattedId, "#"))
	if err != nil {
		return 0, fmt.Errorf("error parsing id from string: %w", err)
	}
	return id, nil
}

func DeformatName(name string) string {
	words := strings.Split(name, " ")
	if len(words) == 1 {
		return fmt.Sprintf("%s%s", strings.ToLower(string(words[0][0])), words[0][1:])
	}

	result := ""
	for _, word := range words {
		lowerCaseWord := strings.ToLower(word)
		result += lowerCaseWord + " "
	}
	return result[:len(result)-1]
}

func FormatHeight(height int32) string {
	feet := int(float64(height) / 3.048)
	inches := int(float64(height) - float64(feet)*3.048)
	return fmt.Sprintf("%d ft %d in", feet, inches)
}

func FormatWeight(weight int32) string {
	return fmt.Sprintf("%.1f lbs", float64(weight)*0.220462)
}

func FormatBaseExp(baseExp int32) string {
	return fmt.Sprintf("%d pts", baseExp)
}

func GetStats(pokemon pokedex.Pokemon) []int {
	return []int{int(pokemon.Hp), int(pokemon.Attack), int(pokemon.Defense), int(pokemon.SpecialAttack), int(pokemon.SpecialDefense), int(pokemon.Speed)}
}

func GetMaxStats() []int {
	// From https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_base_stats_in_Generation_IX
	return []int{255, 190, 250, 194, 250, 200}
}

func GetChatStats(pokemon pokedex.Pokemon) []int {
	return []int{int(pokemon.Hp), int(pokemon.Attack), int(pokemon.Defense), int(pokemon.Speed)}
}

func GetChatMaxStats() []int {
	// From https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_base_stats_in_Generation_IX
	return []int{255, 190, 250, 200}
}
