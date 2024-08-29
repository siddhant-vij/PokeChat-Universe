-- +goose Up
CREATE TABLE pokemons (
  id INTEGER PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  name VARCHAR(255) NOT NULL UNIQUE,
  height INTEGER NOT NULL,
  weight INTEGER NOT NULL,
  picture_url VARCHAR(255) NOT NULL,
  base_experience INTEGER NOT NULL,
  types VARCHAR(255)[] NOT NULL,
  hp INTEGER NOT NULL,
  attack INTEGER NOT NULL,
  defense INTEGER NOT NULL,
  special_attack INTEGER NOT NULL,
  special_defense INTEGER NOT NULL,
  speed INTEGER NOT NULL
);

-- +goose Down
DROP TABLE pokemons;