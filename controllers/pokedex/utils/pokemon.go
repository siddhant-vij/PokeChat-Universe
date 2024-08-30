package utils

import (
	"fmt"
	"strings"
)

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
