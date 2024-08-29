-- name: InsertPokemonEvolution :exec
INSERT INTO evolutions
  (chain_id, pokemon_id, evolves_to_id)
VALUES
  ($1, $2, $3)
ON CONFLICT (chain_id, pokemon_id) DO NOTHING;

-- name: GetFullEvolutionChain :many
WITH simple_chain AS (
  -- Get the given Pokémon and its direct evolution and predecessor
  SELECT 
    p1.id AS pokemon_id,
    p1.name AS pokemon_name,
    pe.evolves_to_id AS evolves_to_id,
    p2.name AS evolves_to_name,
    pe.pokemon_id AS evolves_from_id,
    p3.name AS evolves_from_name
  FROM 
    pokemons p1
  LEFT JOIN 
    evolutions pe ON p1.id = pe.pokemon_id
  LEFT JOIN 
    pokemons p2 ON pe.evolves_to_id = p2.id
  LEFT JOIN 
    evolutions pe2 ON p1.id = pe2.evolves_to_id
  LEFT JOIN 
    pokemons p3 ON pe2.pokemon_id = p3.id
  WHERE 
    p1.id = $1
),
-- Prepare the final results
final_chain AS (
  SELECT 
    pokemon_id AS id,
    pokemon_name AS name,
    1 AS position
  FROM 
    simple_chain
  WHERE 
    pokemon_id IS NOT NULL
  
  UNION ALL
  
  SELECT 
    evolves_from_id AS id,
    evolves_from_name AS name,
    0 AS position
  FROM 
    simple_chain
  WHERE 
    evolves_from_id IS NOT NULL
  
  UNION ALL
  
  SELECT 
    evolves_to_id AS id,
    evolves_to_name AS name,
    2 AS position
  FROM 
    simple_chain
  WHERE 
    evolves_to_id IS NOT NULL
)
-- Final selection and sorting
SELECT 
  id,
  name
FROM 
  final_chain
ORDER BY 
  position;
