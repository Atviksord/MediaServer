-- +goose Up
ALTER TABLE favourites
ADD CONSTRAINT unique_user_media UNIQUE (user_id, media_id);

-- +goose Down
ALTER TABLE favourites
DROP CONSTRAINT unique_user_media;