-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS cards;

CREATE TABLE cards (
    card_id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    word VARCHAR(100) NOT NULL,
    translation VARCHAR(100) NOT NULL,
    easiness FLOAT8 NOT NULL DEFAULT 2.5,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    interval SMALLINT DEFAULT 0,
    expires_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    repetition_number SMALLINT DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cards;
-- +goose StatementEnd