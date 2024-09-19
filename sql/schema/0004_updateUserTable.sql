-- +goose Up
ALTER TABLE users
    ADD COLUMN refreshToken TEXT UNIQUE;

-- +goose Down
ALTER TABLE users
    DROP COLUMN refreshToken;