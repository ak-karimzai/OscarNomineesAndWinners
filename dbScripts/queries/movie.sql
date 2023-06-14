-- name: CreateMovie :one
INSERT INTO "movies" ("title", "release_year", "director", "genre") VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetMovie :one
SELECT * FROM "movies" WHERE id = $1 LIMIT 1;

-- name: ListMovies :many
SELECT * FROM "movies" LIMIT $1 OFFSET $2;

-- name: DeleteMovie :exec
DELETE FROM "movies" WHERE id = $1;

-- name: UpdateMovie :exec
UPDATE "movies" SET title = $1, release_year = $2, director = $3, genre = $4 WHERE id = $5;