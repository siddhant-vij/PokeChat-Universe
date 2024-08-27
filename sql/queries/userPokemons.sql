-- name: InsertUserCollectedPokemon :exec
INSERT INTO user_pokemons
  (id, user_id, pokemon_id)
VALUES
  ($1, $2, $3);

-- name: GetUserCollectedPokemonNames :many
SELECT pokemons.name
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
WHERE user_id = $1;

-- name: GetUserAvailablePokemonNames :many
SELECT pokemons.name
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
WHERE user_id = $1 AND user_pokemons.id IS NULL;

-- name: GetUserCollectedPokemonsSortedByIdAsc :many
SELECT pokemons.id, pokemons.name, pokemons.picture_url
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
WHERE user_id = $1
ORDER BY pokemons.id ASC
LIMIT $2 OFFSET $3;

-- name: GetUserAvailablePokemonsSortedByIdAsc :many
SELECT pokemons.id, pokemons.name, pokemons.picture_url
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
WHERE user_id = $1 AND user_pokemons.id IS NULL
ORDER BY pokemons.id ASC
LIMIT $2 OFFSET $3;

-- name: GetUserCollectedPokemonsSortedByIdDesc :many
SELECT pokemons.id, pokemons.name, pokemons.picture_url
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
WHERE user_id = $1
ORDER BY pokemons.id DESC
LIMIT $2 OFFSET $3;

-- name: GetUserAvailablePokemonsSortedByIdDesc :many
SELECT pokemons.id, pokemons.name, pokemons.picture_url
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
WHERE user_id = $1 AND user_pokemons.id IS NULL
ORDER BY pokemons.id DESC
LIMIT $2 OFFSET $3;

-- name: GetUserCollectedPokemonsSortedByNameAsc :many
SELECT pokemons.id, pokemons.name, pokemons.picture_url
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
WHERE user_id = $1
ORDER BY pokemons.name ASC
LIMIT $2 OFFSET $3;

-- name: GetUserAvailablePokemonsSortedByNameAsc :many
SELECT pokemons.id, pokemons.name, pokemons.picture_url
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
WHERE user_id = $1 AND user_pokemons.id IS NULL
ORDER BY pokemons.name ASC
LIMIT $2 OFFSET $3;

-- name: GetUserCollectedPokemonsSortedByNameDesc :many
SELECT pokemons.id, pokemons.name, pokemons.picture_url
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
WHERE user_id = $1
ORDER BY pokemons.name DESC
LIMIT $2 OFFSET $3;

-- name: GetUserAvailablePokemonsSortedByNameDesc :many
SELECT pokemons.id, pokemons.name, pokemons.picture_url
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
WHERE user_id = $1 AND user_pokemons.id IS NULL
ORDER BY pokemons.name DESC
LIMIT $2 OFFSET $3;