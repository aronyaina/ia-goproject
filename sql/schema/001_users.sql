-- +goose Up
CREATE TABLE users (
  user_id UUID PRIMARY KEY,
  user_email STRING(25),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users