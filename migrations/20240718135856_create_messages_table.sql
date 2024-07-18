-- +goose Up
-- +goose StatementBegin
create table IF NOT EXISTS messages (
    id                      BIGSERIAL PRIMARY KEY,
    chat_id                 BIGSERIAL REFERENCES chats(id) NOT NULL,
    user_id                 BIGSERIAL REFERENCES users(id) NOT NULL,
    message                 VARCHAR(300) NOT NULL,
    created_at               TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table messages;
-- +goose StatementEnd
