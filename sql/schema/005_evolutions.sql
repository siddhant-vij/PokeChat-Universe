-- +goose Up
CREATE TABLE evolutions (
  id SERIAL PRIMARY KEY,
  chain_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  pokemon_id INTEGER NOT NULL REFERENCES pokemons(id) ON DELETE CASCADE,
  evolves_to_id INTEGER REFERENCES pokemons(id) ON DELETE CASCADE,
  UNIQUE (chain_id, pokemon_id)
);

-- +goose Down
DROP TABLE evolutions;