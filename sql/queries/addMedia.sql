-- name: AddMedia :one
INSERT INTO media (media_name, media_type, file_path, format, upload_date, follow_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
