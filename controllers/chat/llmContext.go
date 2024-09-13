package chat

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/pokedex/utils"
)

// number of recent messages to set up chat context
var recentMessages = 4

func GenerateMessage(currentUserMessage, pokemonName string, cfg *config.AppConfig) []api.Message {
	currentPrompt := fmt.Sprintf(`I want you to act like Pokemon: %s, staying true to its nature, characteristics, and abilities in all of its forms - to response to the below mentioned user's message. Respond as if you are the Pokemon itself, keeping the tone friendly and engaging, but concise. Provide responses that are brief, clear, and to the point. Avoid overly complex language. Keep the conversation fun and easy to understand. Answer the user's questions thoughtfully, while maintaining the unique personality of the Pokemon.
	
User's message: %s`, pokemonName, currentUserMessage)

	pokemon, err := cfg.DBQueries.GetPokemonDetailsByName(context.Background(), utils.DeformatName(pokemonName))
	if err != nil {
		log.Println(err)
		return nil
	}

	chatMsgParams := pokedex.GetAllChatsForUserAndPokemonParams{
		UserID:    cfg.LoggedInUserId,
		PokemonID: int32(pokemon.ID),
	}
	chatMsgs, err := cfg.DBQueries.GetAllChatsForUserAndPokemon(context.Background(), chatMsgParams)
	if err != nil {
		log.Println(err)
		return nil
	}

	var orderedChatMessages []ChatMessage
	for _, chatMsg := range chatMsgs {
		orderedChatMessages = append(orderedChatMessages, ChatMessage{
			Sender:  chatMsg.Sender,
			Message: chatMsg.Message,
		})
	}

	length := len(orderedChatMessages)
	if length > recentMessages+1 {
		orderedChatMessages = orderedChatMessages[length-recentMessages-1:]
	}

	messages := []api.Message{}
	length = len(orderedChatMessages)
	for i := 0; i < length; i++ {
		if i == length-1 {
			continue
		}
		messages = append(messages, api.Message{
			Role:    orderedChatMessages[i].Sender,
			Content: orderedChatMessages[i].Message,
		})
	}

	messages = append(messages, api.Message{
		Role:    "user",
		Content: currentPrompt,
	})
	return messages
}
