-- +goose Up
-- +goose StatementBegin
create table IF NOT EXISTS chat_users (
    chat_id                 BIGSERIAL REFERENCES chats(id) NOT NULL,
    user_id                 BIGSERIAL REFERENCES users(id) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chat_users;
-- +goose StatementEnd
