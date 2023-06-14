-- name: CreatePerformance :one
INSERT INTO performances(actor_id, movie_id, year) VALUES ($1, $2, $3) RETURNING *;

-- name: GetPerformance :one
SELECT * FROM performances WHERE id = $1;

-- name: ListPerformances :many
SELECT * FROM performances LIMIT $1 OFFSET  $2;

-- name: UpdatePerformance :exec
UPDATE performances SET actor_id = $2, movie_id = $3, year = $4 WHERE id = $1;