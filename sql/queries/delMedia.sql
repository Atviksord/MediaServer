-- name: DeleteMedia :one
DELETE FROM media
WHERE id = $1
RETURNING *;
