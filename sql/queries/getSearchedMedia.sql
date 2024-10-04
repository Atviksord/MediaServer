-- name: GetSearchedMedia :many
SELECT *
FROM media
WHERE media_name LIKE '%' || $1 || '%';
