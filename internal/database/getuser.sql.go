// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: getuser.sql

package database

import (
	"context"
)

const getUser = `-- name: GetUser :one
SELECT id, username, password, created_at, updated_at, refreshtoken FROM users WHERE username = $1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Refreshtoken,
	)
	return i, err
}
