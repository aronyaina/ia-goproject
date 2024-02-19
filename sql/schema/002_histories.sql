-- +goose Up
CREATE TABLE histories (
  history_id UUID PRIMARY KEY,
  history_title TEXT NOT NULL,
  user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
  prompt_id UUID REFERENCES prompts(prompt_id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE histories;