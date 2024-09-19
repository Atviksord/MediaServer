-- +goose Up
-- Create a table for users' favourite media
CREATE TABLE favourites (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,  -- Foreign key to users
    media_id INTEGER NOT NULL REFERENCES media(id) ON DELETE CASCADE, -- Foreign key to media
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- When the media was favorited
);

-- Create a table for users following media (e.g., subscribing to updates or similar)
CREATE TABLE follows (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,  -- Foreign key to users
    media_id INTEGER NOT NULL REFERENCES media(id) ON DELETE CASCADE, -- Foreign key to media
    followed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- When the follow was initiated
);

-- +goose Down
-- Drop the favourites and follows tables
DROP TABLE favourites;
DROP TABLE follows;
