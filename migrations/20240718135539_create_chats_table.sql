-- +goose Up
-- +goose StatementBegin
create table IF NOT EXISTS chats (
    id                      BIGSERIAL PRIMARY KEY,
    name                    VARCHAR(50) NOT NULL UNIQUE,
    is_active               BOOL NOT NULL DEFAULT true
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chats;
-- +goose StatementEnd
