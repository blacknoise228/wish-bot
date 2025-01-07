-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
  "username" varchar UNIQUE NOT NULL,
  "chat_id" bigint NOT NULL PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "dim_wish_status" (
    id integer NOT NULL PRIMARY KEY,
    status_name varchar NOT NULL
);

INSERT INTO "dim_wish_status" (id, status_name) VALUES (1, 'public');

INSERT INTO "dim_wish_status" (id, status_name) VALUES (2, 'only friends');

CREATE TABLE "wish" (
    id serial NOT NULL PRIMARY KEY,
    chat_id bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    description varchar NOT NULL,
    link varchar NOT NULL,
    status integer NOT NULL
);

CREATE TABLE "dim_friend_status" (
    id integer NOT NULL PRIMARY KEY,
    status_name varchar NOT NULL
);

INSERT INTO "dim_friend_status" (id, status_name) VALUES (1, 'approved');

INSERT INTO "dim_friend_status" (id, status_name) VALUES (2, 'pending');

INSERT INTO "dim_friend_status" (id, status_name) VALUES (3, 'declined');

CREATE TABLE "friends" (
    chat_id bigint NOT NULL,
    friend_id bigint NOT NULL,
    status integer NOT NULL DEFAULT 2,
    created_at timestamptz NOT NULL DEFAULT (now()),
    PRIMARY KEY (chat_id, friend_id)
);

ALTER TABLE "wish" ADD FOREIGN KEY ("chat_id") REFERENCES "users" ("chat_id");

ALTER TABLE "wish" ADD FOREIGN KEY ("status") REFERENCES "dim_wish_status" ("id");

ALTER TABLE "friends" ADD FOREIGN KEY ("chat_id") REFERENCES "users" ("chat_id");

ALTER TABLE "friends" ADD FOREIGN KEY ("friend_id") REFERENCES "users" ("chat_id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists friends;

drop table if exists wish;

drop table if exists dim_friend_status;

drop table if exists dim_wish_status;

drop table if exists users;
-- +goose StatementEnd
