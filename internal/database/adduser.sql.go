// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: adduser.sql

package database

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(username,password,created_at,updated_at)
VALUES($1,$2,$3,$4)
RETURNING id, username, password, created_at, updated_at, refreshtoken
`

type CreateUserParams struct {
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
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
