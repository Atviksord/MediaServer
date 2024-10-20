-- name: GetFavouritedMedia :one
SELECT COUNT(*)
FROM favourites
WHERE user_id = $1 AND media_id = $2;
