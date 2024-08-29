-- +goose Up
CREATE INDEX pokemons_name_idx ON pokemons(name);

-- +goose Down
DROP INDEX pokemons_name_idx;