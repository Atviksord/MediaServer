-- name: GetMedia :many
SELECT media.*
FROM media
JOIN favourites ON media.id = favourites.media_id
WHERE favourites.user_id = $1; 
