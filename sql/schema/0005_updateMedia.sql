-- +goose Up
ALTER TABLE media
    DROP COLUMN file_size,
    DROP COLUMN duration,
    DROP COLUMN resolution,
    DROP COLUMN thumbnail,
    DROP COLUMN tags,
    DROP COLUMN description,
    DROP COLUMN view_count,
    DROP COLUMN rating;

-- +goose Down
ALTER TABLE media
    ADD COLUMN file_size BIGINT,
    ADD COLUMN duration INTEGER,
    ADD COLUMN resolution TEXT,
    ADD COLUMN thumbnail TEXT,
    ADD COLUMN tags TEXT[],
    ADD COLUMN description TEXT,
    ADD COLUMN view_count INTEGER DEFAULT 0,
    ADD COLUMN rating DECIMAL(2,1);