-- name: CreatePlayer :one
INSERT INTO player(name,gender,age,total_pizza,logged_pizza,coins,date_created) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING *;
-- name: GetPlayers :many
SELECT * FROM player;
-- name: GetPlayer :one
SELECT * FROM player WHERE id = $1;
-- name: DeletePlayer :exec
DELETE FROM player WHERE id = $1;
-- name: UpdatePlayer :one
UPDATE player
SET name = $1,gender = $2,age = $3,total_pizza = $4,logged_pizza = $5,coins = $6,date_updated = CURRENT_TIMESTAMP
WHERE id = $7
RETURNING *;
 
