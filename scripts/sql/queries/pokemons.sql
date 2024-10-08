-- name: InsertPokemon :exec
INSERT INTO pokemons
  (id, name, height, weight, picture_url, base_experience, types, hp, attack, defense, special_attack, special_defense, speed)
VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);

-- name: GetPokemonCount :one
SELECT COUNT(*) FROM pokemons;

-- name: GetPokemonDetailsByName :one
SELECT * FROM pokemons
WHERE name = $1 LIMIT 1;

-- name: GetPokemonByID :one
SELECT * FROM pokemons
WHERE id = $1 LIMIT 1;

-- name: UpdatePokemonByID :exec
UPDATE pokemons
SET
  name = $2,
  height = $3,
  weight = $4,
  picture_url = $5,
  base_experience = $6,
  types = $7,
  hp = $8,
  attack = $9,
  defense = $10,
  special_attack = $11,
  special_defense = $12,
  speed = $13,
  created_at = $14,
  updated_at = NOW()
WHERE id = $1;

-- name: DeletePokemonByID :exec
DELETE FROM pokemons
WHERE id = $1;

-- name: SearchPokemonByName :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
WHERE pokemons.name ILIKE $1
LIMIT $2;

-- name: GetPokemonsSortedByIdAsc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
ORDER BY pokemons.id ASC
LIMIT $1 OFFSET $2;

-- name: GetPokemonsSortedByIdDesc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
ORDER BY pokemons.id DESC
LIMIT $1 OFFSET $2;

-- name: GetPokemonsSortedByNameAsc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
ORDER BY pokemons.name ASC
LIMIT $1 OFFSET $2;

-- name: GetPokemonsSortedByNameDesc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
ORDER BY pokemons.name DESC
LIMIT $1 OFFSET $2;