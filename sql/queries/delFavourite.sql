-- name: DeleteFavourite :one
DELETE FROM favourites
WHERE user_id = $1 AND media_id = $2
RETURNING *;
