-- name: GetUserByToken :one
SELECT * FROM users WHERE  refreshToken = $1;