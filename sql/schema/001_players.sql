-- +goose Up
CREATE TABLE player (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50),
  gender VARCHAR(25),
  age INTEGER,
  total_pizza INTEGER,
  logged_pizza INTEGER,
  coins INTEGER,
  date_created TIMESTAMP
);
-- +goose Down
DROP TABLE player;