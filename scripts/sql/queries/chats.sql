-- name: InsertChatEntry :exec
INSERT INTO chats
  (id, user_pokemon_id)
SELECT
  $1,
  up.id
FROM user_pokemons up
WHERE up.user_id = $2
  AND up.pokemon_id = $3
ON CONFLICT (user_pokemon_id) DO UPDATE SET updated_at = NOW();
