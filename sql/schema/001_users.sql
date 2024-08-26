-- +goose Up
CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  auth_id VARCHAR(255) NOT NULL UNIQUE,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  picture_url VARCHAR(255)
);

-- +goose Down
DROP TABLE users;