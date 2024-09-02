package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func FormatID(id int) string {
	return fmt.Sprintf("#%04d", id)
}

func DisplayTypes(types []string) []string {
	if len(types) <= 2 {
		result := make([]string, 0)
		for _, pokemonType := range types {
			pokemonType = FormatName(pokemonType)
			result = append(result, pokemonType)
		}
		return result
	}

	indices := make([]int, 2)
	idx0, _ := rand.Int(rand.Reader, big.NewInt(int64(len(types))))
	indices[0] = int(idx0.Int64())
	idx1, _ := rand.Int(rand.Reader, big.NewInt(int64(len(types))))
	indices[1] = int(idx1.Int64())
	for indices[0] == indices[1] {
		idx1, _ := rand.Int(rand.Reader, big.NewInt(int64(len(types))))
		indices[1] = int(idx1.Int64())
	}

	return []string{
		FormatName(types[indices[0]]),
		FormatName(types[indices[1]]),
	}
}

func FormatName(name string) string {
	words := strings.Split(name, " ")
	if len(words) == 1 {
		return fmt.Sprintf("%s%s", strings.ToUpper(string(words[0][0])), words[0][1:])
	}

	result := ""
	for _, word := range words {
		firstChar := strings.ToUpper(string(word[0]))
		word = firstChar + word[1:]
		result += word + " "
	}
	return result[:len(result)-1]
}
