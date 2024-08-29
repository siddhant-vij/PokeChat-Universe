-- +goose Up
CREATE TABLE user_pokemons (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  pokemon_id INTEGER NOT NULL REFERENCES pokemons(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE user_pokemons;