-- +goose Up
CREATE TABLE media (
    id SERIAL PRIMARY KEY,
    media_name TEXT NOT NULL,
    media_type TEXT NOT NULL,  -- e.g., movie, image, audio, etc.
    file_path TEXT NOT NULL,   -- path where the media is stored
    file_size BIGINT,          -- file size in bytes
    duration INTEGER,          -- duration in seconds (for videos/audio)
    resolution TEXT,           -- e.g., 1920x1080 for videos or images
    format TEXT NOT NULL,      -- file format, e.g., mp4, jpg, mp3
    upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- time of upload
    thumbnail TEXT,            -- reference to a thumbnail image
    tags TEXT[],               -- array of tags for categorizing media
    description TEXT,          -- brief description or synopsis
    view_count INTEGER DEFAULT 0,  -- track number of views
    rating DECIMAL(2,1),       -- optional rating, e.g., 4.5 out of 5
    follow_id INTEGER          -- foreign key reference (optional, can be used to group media)
);

-- +goose Down
DROP TABLE media;