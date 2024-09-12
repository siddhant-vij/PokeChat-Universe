-- name: InsertChatHistory :exec
INSERT INTO chat_histories
  (id, chat_id, sender, message)
SELECT
  $1,
  c.id,
  $4,
  $5
FROM user_pokemons up
JOIN chats c ON up.id = c.user_pokemon_id
WHERE up.user_id = $2
  AND up.pokemon_id = $3;

-- name: GetAllChatsForUserAndPokemon :many
SELECT
  ch.sender,
  ch.message
FROM chat_histories ch
INNER JOIN chats c ON ch.chat_id = c.id
INNER JOIN user_pokemons up ON c.user_pokemon_id = up.id
WHERE up.user_id = $1
  AND up.pokemon_id = $2
ORDER BY ch.created_at ASC;