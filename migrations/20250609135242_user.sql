-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS users (
        user_id VARCHAR(36) PRIMARY KEY NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        hashed_password VARCHAR(255) NOT NULL,
        registration_date TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;

-- +goose StatementEnd