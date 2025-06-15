-- +goose Up
-- +goose StatementBegin

CREATE TABLE results (
    result_id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    card_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    grade INTEGER NOT NULL,

    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT fk_card
        FOREIGN KEY (card_id)
        REFERENCES cards(card_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS results;
-- +goose StatementEnd

