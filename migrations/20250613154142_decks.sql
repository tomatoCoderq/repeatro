-- +goose Up
-- +goose StatementBegin

DROP TABLE IF EXISTS decks;

CREATE TABLE decks (
    deck_id UUID PRIMARY KEY,
    created_by UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    name VARCHAR(100) NOT NULL,
    description VARCHAR(100),

    CONSTRAINT fk_user_created_by
        FOREIGN KEY (created_by)
        REFERENCES users(user_id)
        ON DELETE CASCADE
);

CREATE INDEX idx_decks_created_by ON decks (created_by);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS decks;
-- +goose StatementEnd
