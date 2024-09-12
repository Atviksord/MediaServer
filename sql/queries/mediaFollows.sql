-- name: GetFollowedMedia :many
SELECT media.*
FROM media
JOIN follows ON media.id = follows.media_id
WHERE follows.user_id = 1;  -- Replace with the desired user ID
