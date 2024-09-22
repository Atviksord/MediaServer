-- name: DelAccessToken :one
UPDATE users
SET refreshToken = NULL
WHERE refreshToken = $1
RETURNING *;