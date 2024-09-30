-- name: DeleteMedia :one
DELETE FROM media
WHERE file_path = $1
RETURNING *;
