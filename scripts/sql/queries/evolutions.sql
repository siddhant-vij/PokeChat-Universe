-- name: InsertPokemonEvolution :exec
INSERT INTO evolutions
  (chain_id, pokemon_id, evolves_to_id)
VALUES
  ($1, $2, $3)
ON CONFLICT (chain_id, pokemon_id) DO NOTHING;

-- name: GetFullEvolutionChain :many
WITH given_pokemon AS (
  -- Select the given Pokémon
  SELECT
    p.id AS pokemon_id,
    p.name AS pokemon_name,
    p.picture_url AS picture_url,
    1 AS position
  FROM
    pokemons p
  WHERE
    p.id = $1
),

predecessor AS (
  -- Find the predecessor Pokémon if it exists
  SELECT
    p2.id AS pokemon_id,
    p2.name AS pokemon_name,
    p2.picture_url AS picture_url,
    0 AS position
  FROM
    evolutions e
  JOIN
    pokemons p1 ON e.evolves_to_id = p1.id
  JOIN
    pokemons p2 ON e.pokemon_id = p2.id
  WHERE
    p1.id = $1
),

successor AS (
  -- Find the successor Pokémon if it exists
  SELECT
    p2.id AS pokemon_id,
    p2.name AS pokemon_name,
    p2.picture_url AS picture_url,
    2 AS position
  FROM
    evolutions e
  JOIN
    pokemons p1 ON e.pokemon_id = p1.id
  JOIN
    pokemons p2 ON e.evolves_to_id = p2.id
  WHERE
    p1.id = $1
)

-- Combine results and sort by position
SELECT
  pokemon_id AS id,
  pokemon_name AS name,
  picture_url
FROM (
  SELECT
    pokemon_id,
    pokemon_name,
    picture_url,
    position
  FROM
    predecessor
  
  UNION ALL
  
  SELECT
    pokemon_id,
    pokemon_name,
    picture_url,
    position
  FROM
    given_pokemon
  
  UNION ALL
  
  SELECT
    pokemon_id,
    pokemon_name,
    picture_url,
    position
  FROM
    successor
) combined
ORDER BY
  position DESC;