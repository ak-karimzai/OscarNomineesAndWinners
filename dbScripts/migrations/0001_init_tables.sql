-- +goose Up

CREATE TABLE "movies" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "release_year" integer NOT NULL,
  "director" varchar NOT NULL,
  "genre" varchar NOT NULL
);

CREATE TABLE "actors" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "birth_year" integer NOT NULL,
  "nationality" varchar NOT NULL
);

CREATE TABLE "awards" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "category" varchar NOT NULL
);

CREATE TABLE "nominations" (
  "id" bigserial PRIMARY KEY,
  "movie_id" int NOT NULL,
  "award_id" int NOT NULL,
  "year" integer NOT NULL,
  "is_winner" boolean NOT NULL
);

CREATE TABLE "performances" (
  "id" bigserial PRIMARY KEY,
  "actor_id" int NOT NULL,
  "movie_id" int NOT NULL,
  "year" int NOT NULL
);

CREATE TABLE "nominated_performances" (
  "nomination_id" int,
  "performance_id" int,
  PRIMARY KEY ("nomination_id", "performance_id")
);

ALTER TABLE "nominations" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("id");

ALTER TABLE "nominations" ADD FOREIGN KEY ("award_id") REFERENCES "awards" ("id");

ALTER TABLE "performances" ADD FOREIGN KEY ("actor_id") REFERENCES "actors" ("id");

ALTER TABLE "performances" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("id");

ALTER TABLE "nominated_performances" ADD FOREIGN KEY ("nomination_id") REFERENCES "nominations" ("id");

ALTER TABLE "nominated_performances" ADD FOREIGN KEY ("performance_id") REFERENCES "performances" ("id");

-- +goose Down

DROP TABLE "nominated_performances";
DROP TABLE "performances";
DROP TABLE "nominations";
DROP TABLE "awards";
DROP TABLE "movies";
DROP TABLE "actors";
