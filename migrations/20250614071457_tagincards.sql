-- +goose Up
-- +goose StatementBegin
-- CREATE TABLE IF NOT EXISTS cards;

ALTER TABLE cards ADD COLUMN tags TEXT[];

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE cards DROP COLUMN tags;

-- +goose StatementEnd
