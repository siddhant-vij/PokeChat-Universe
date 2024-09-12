-- +goose Up
CREATE TABLE chat_histories (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  chat_id UUID NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
  sender VARCHAR(20) NOT NULL,
  message TEXT NOT NULL
);

-- +goose Down
DROP TABLE chat_histories;