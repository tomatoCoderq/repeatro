-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS cards;
CREATE TABLE IF NOT EXISTS cards (
    card_id VARCHAR(36) PRIMARY KEY NOT NULL,
    word VARCHAR(100) NOT NULL,
    translation VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    repetition_number SMALLINT NOT NULL DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cards;
-- +goose StatementEnd