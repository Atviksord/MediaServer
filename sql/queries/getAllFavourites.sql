-- name: GetAllFavouriteMedia :many
SELECT media.id, media.media_name, media.media_type, media.file_path, media.format, media.upload_date, media.follow_id
FROM media
JOIN favourites ON media.id = favourites.media_id
WHERE favourites.user_id = $1;
