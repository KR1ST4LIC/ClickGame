-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA clans;

CREATE TABLE clans.clan
(
    "id"      serial PRIMARY KEY,
    "name"    text
);

CREATE TABLE clans.user_clan
(
    "user_id" int references users.user (id),
    "clan_id" int references clans.clan (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clans.user_clan;
DROP TABLE IF EXISTS clans.clan;
DROP SCHEMA clans CASCADE;
-- +goose StatementEnd
