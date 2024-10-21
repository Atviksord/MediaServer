-- name: AddFavourite :one
INSERT INTO favourites(user_id,media_id)
VALUES($1,$2)
RETURNING *;