-- +goose Up
CREATE INDEX user_pokemons_user_idx ON user_pokemons(user_id);
CREATE INDEX user_pokemons_pokemon_idx ON user_pokemons(pokemon_id);
CREATE INDEX chat_histories_chat_idx ON chat_histories(chat_id);
CREATE INDEX chat_histories_created_at_idx ON chat_histories(created_at);

-- +goose Down
DROP INDEX chat_histories_created_at_idx;
DROP INDEX chat_histories_chat_idx;
DROP INDEX user_pokemons_pokemon_idx;
DROP INDEX user_pokemons_user_idx;