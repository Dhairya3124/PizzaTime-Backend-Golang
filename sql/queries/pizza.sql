-- name: CreatePizza :one
INSERT INTO pizza(player_id,logged_pizza,date_created) VALUES($1,$2,$3) RETURNING *;

