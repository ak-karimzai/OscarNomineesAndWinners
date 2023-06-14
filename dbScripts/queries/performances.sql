-- name: CreatePerformance :one
INSERT INTO performances(actor_id, movie_id, year) VALUES ($1, $2, $3) RETURNING *;

-- name: GetPerformance :one
SELECT * FROM performances WHERE id = $1;
