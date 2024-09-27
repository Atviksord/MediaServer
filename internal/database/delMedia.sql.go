// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: delMedia.sql

package database

import (
	"context"
)

const deleteMedia = `-- name: DeleteMedia :one
DELETE FROM media
WHERE id = $1
RETURNING id, media_name, media_type, file_path, format, upload_date, follow_id
`

func (q *Queries) DeleteMedia(ctx context.Context, id int32) (Medium, error) {
	row := q.db.QueryRowContext(ctx, deleteMedia, id)
	var i Medium
	err := row.Scan(
		&i.ID,
		&i.MediaName,
		&i.MediaType,
		&i.FilePath,
		&i.Format,
		&i.UploadDate,
		&i.FollowID,
	)
	return i, err
}
