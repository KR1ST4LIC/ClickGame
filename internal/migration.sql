CREATE TABLE clans
(
    "id"      serial PRIMARY KEY,
    "name"    text,
    "balance" bigint,
    "top"     int
);
CREATE TABLE users
(
    "id"         serial PRIMARY KEY,
    "tg_id"      bigint UNIQUE,
    "name"       text,
    "balance"    bigint default 0,
    "multiplier" real default 1.0,
    "click"      int default 1,
    "afk"        int default 0,
    "barrier"    int default 1000,
    "clan"       int references clans (id),
    "status"     text default user
);
CREATE TABLE top
(
    "id"      int references users (id),
    "balance" bigint,
    "top"     int
);
CREATE TABLE bestclans
(
    "id"      int references clans (id),
    "balance" bigint,
    "top"     int
);