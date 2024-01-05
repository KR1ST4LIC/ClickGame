-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA users;

CREATE TABLE users.user
(
    "id"          serial PRIMARY KEY,
    "tg_id"       bigint UNIQUE,
    "name"        text,
    "balance"     bigint default 0 CHECK ( balance >= 0 ) CHECK (balance <= max_balance),
    "multiplier"  real   default 1,
    "click"       int    default 1,
    "afk"         int    default 0,
    "max_balance" bigint default 1000,
    "status"      text   default user
);
CREATE INDEX balance_index ON users.user (balance);
CREATE TABLE users.sec_user
(
    "tg_id" bigint UNIQUE,
    "addbal" bigint default 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users.sec_user;
DROP TABLE IF EXISTS users.user;

DROP SCHEMA users CASCADE;
-- +goose StatementEnd
