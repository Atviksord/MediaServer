// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: getSearchedMedia.sql

package database

import (
	"context"
	"database/sql"
)

const getSearchedMedia = `-- name: GetSearchedMedia :many
SELECT id, media_name, media_type, file_path, format, upload_date, follow_id
FROM media
WHERE media_name LIKE '%' || $1 || '%'
`

func (q *Queries) GetSearchedMedia(ctx context.Context, dollar_1 sql.NullString) ([]Medium, error) {
	rows, err := q.db.QueryContext(ctx, getSearchedMedia, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Medium
	for rows.Next() {
		var i Medium
		if err := rows.Scan(
			&i.ID,
			&i.MediaName,
			&i.MediaType,
			&i.FilePath,
			&i.Format,
			&i.UploadDate,
			&i.FollowID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
