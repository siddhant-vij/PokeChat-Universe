-- +goose Up
CREATE INDEX users_auth_id_idx ON users (auth_id);

-- +goose Down
DROP INDEX users_auth_id_idx;