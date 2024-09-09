-- +goose Up
CREATE TABLE chats (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  user_pokemon_id UUID NOT NULL UNIQUE REFERENCES user_pokemons(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE chats;