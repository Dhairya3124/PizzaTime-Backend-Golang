-- name: CreatePlayer :one
INSERT INTO player(name,gender,age,total_pizza,logged_pizza,coins,date_created) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING *;
-- name: GetPlayers :many
SELECT * FROM player;
-- name: GetPlayer :one
SELECT * FROM player WHERE id = $1;
