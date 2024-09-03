-- name: InsertUserCollectedPokemon :exec
INSERT INTO user_pokemons
  (id, user_id, pokemon_id)
VALUES
  ($1, $2, $3);

-- name: IsPokemonCollected :one
SELECT
  EXISTS (
    SELECT 1
    FROM user_pokemons
    WHERE user_id = $1 AND pokemon_id = $2
  );

-- name: SearchUserCollectedPokemonsByName :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE pokemons.name ILIKE $2
ORDER BY pokemons.name ASC
LIMIT $3;

-- name: SearchUserAvailablePokemonsByName :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
AND pokemons.name ILIKE $2
ORDER BY pokemons.name ASC
LIMIT $3;

-- name: GetUserCollectedPokemonsSortedByIdAsc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NOT NULL
ORDER BY pokemons.id ASC
LIMIT $2 OFFSET $3;

-- name: GetUserAvailablePokemonsSortedByIdAsc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
ORDER BY pokemons.id ASC
LIMIT $2 OFFSET $3;

-- name: GetUserCollectedPokemonsSortedByIdDesc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NOT NULL
ORDER BY pokemons.id DESC
LIMIT $2 OFFSET $3;

-- name: GetUserAvailablePokemonsSortedByIdDesc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
ORDER BY pokemons.id DESC
LIMIT $2 OFFSET $3;

-- name: GetUserCollectedPokemonsSortedByNameAsc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NOT NULL
ORDER BY pokemons.name ASC
LIMIT $2 OFFSET $3;

-- name: GetUserAvailablePokemonsSortedByNameAsc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
ORDER BY pokemons.name ASC
LIMIT $2 OFFSET $3;

-- name: GetUserCollectedPokemonsSortedByNameDesc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NOT NULL
ORDER BY pokemons.name DESC
LIMIT $2 OFFSET $3;

-- name: GetUserAvailablePokemonsSortedByNameDesc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
ORDER BY pokemons.name DESC
LIMIT $2 OFFSET $3;

-- name: GetOneAvailablePokemonAfterCollectionByIdAsc :one
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
AND pokemons.id > $2
ORDER BY pokemons.id ASC
LIMIT 1;

-- name: GetOneAvailablePokemonAfterCollectionByIdDesc :one
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
AND pokemons.id < $2
ORDER BY pokemons.id DESC
LIMIT 1;

-- name: GetOneAvailablePokemonAfterCollectionByNameAsc :one
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
AND pokemons.name > $2
ORDER BY pokemons.name ASC
LIMIT 1;

-- name: GetOneAvailablePokemonAfterCollectionByNameDesc :one
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
AND pokemons.name < $2
ORDER BY pokemons.name DESC
LIMIT 1;