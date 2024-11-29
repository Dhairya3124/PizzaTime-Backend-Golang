-- +goose Up
CREATE TABLE pizza (
  id SERIAL PRIMARY KEY,
  player_id INTEGER REFERENCES player(id) ON DELETE CASCADE,
  logged_pizza INTEGER NOT NULL,
  date_created TIMESTAMP NOT NULL
);
-- +goose Down
DROP TABLE pizza;