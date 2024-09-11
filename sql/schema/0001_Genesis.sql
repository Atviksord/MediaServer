-- +goose Up
CREATE TABLE users(id SERIAL PRIMARY KEY, 
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL);

-- +goose Down
DROP TABLE users;