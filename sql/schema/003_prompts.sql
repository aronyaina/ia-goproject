-- +goose Up
CREATE TABLE prompts (
  prompt_id UUID PRIMARY KEY,
  prompt_tag TEXT UNIQUE NOT NULL, 
  prompt_results TEXT, 
  user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE prompts;