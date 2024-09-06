-- name: CreateUser :one
INSERT INTO users(username,password,created_at,updated_at)
VALUES($1,$2,$3,$4)
RETURNING *;