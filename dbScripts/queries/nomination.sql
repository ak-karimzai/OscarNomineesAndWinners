-- name: CreateNomination :one
INSERT INTO nominations("movie_id", "award_id", "year", "is_winner") VALUES ($1, $2, $3, $4) RETURNING *;
-- 
-- name: GetNomination :one
SELECT * FROM nominations WHERE id = $1 LIMIT 1;
-- 
-- name: ListNominations :many
SELECT * FROM nominations LIMIT $1 OFFSET $2;
-- 
-- name: DeleteNomination :exec
DELETE FROM nominations WHERE id = $1;
-- 
-- name: UpdateNomination :exec
UPDATE nominations SET movie_id = $1, award_id = $2, year = $3, is_winner = $4 WHERE id = $5;