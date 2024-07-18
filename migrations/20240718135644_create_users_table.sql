-- +goose Up
-- +goose StatementBegin
create table IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    login varchar(50) UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
