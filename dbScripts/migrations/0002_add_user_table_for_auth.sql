-- +goose Up
CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "password" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_last_changed" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

INSERT INTO "users"("username", "password", "fullname", "email") values ('admin', 'admin', 'Muhammad Feroz Nawrozi', 'mfnawrozi@gmail.com');

-- +goose Down
DROP TABLE "users";