// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: addFavourite.sql

package database

import (
	"context"
)

const addFavourite = `-- name: AddFavourite :one
INSERT INTO favourites(user_id,media_id)
VALUES($1,$2)
RETURNING id, user_id, media_id, added_at
`

type AddFavouriteParams struct {
	UserID  int32
	MediaID int32
}

func (q *Queries) AddFavourite(ctx context.Context, arg AddFavouriteParams) (Favourite, error) {
	row := q.db.QueryRowContext(ctx, addFavourite, arg.UserID, arg.MediaID)
	var i Favourite
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.MediaID,
		&i.AddedAt,
	)
	return i, err
}
