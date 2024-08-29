-- +goose Up
CREATE INDEX evolutions_pokemon_id_idx ON evolutions(pokemon_id);
CREATE INDEX evolutions_evolves_to_id_idx ON evolutions(evolves_to_id);

-- +goose Down
DROP INDEX evolutions_evolves_to_id_idx;
DROP INDEX evolutions_pokemon_id_idx;