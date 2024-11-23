-- name: CreatePlayer :one
INSERT INTO player(name,gender,age,total_pizza,logged_pizza,coins,date_created) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING *;