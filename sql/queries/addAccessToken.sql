-- name: AddAccessToken :one
UPDATE users
SET refreshToken = $2
WHERE username = $1
RETURNING *;