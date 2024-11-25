-- +goose Up
CREATE TABLE player (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  gender VARCHAR(25) NOT NULL,
  age INTEGER NOT NULL,
  total_pizza INTEGER NOT NULL DEFAULT 0,
  logged_pizza INTEGER NOT NULL DEFAULT 0,
  coins INTEGER NOT NULL DEFAULT 500,
  date_created TIMESTAMP NOT NULL
);
-- +goose Down
DROP TABLE player;