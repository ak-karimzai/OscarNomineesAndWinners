-- name: CreateAward :one
INSERT INTO awards("name", "category") VALUES ($1, $2) RETURNING id;
-- 
-- name: GetAward :one
SELECT * FROM awards WHERE id = $1 LIMIT 1;
-- 
-- name: ListAwards :many
SELECT * FROM awards LIMIT $1 OFFSET $2;
-- 
-- name: DeleteAward :exec
DELETE FROM awards WHERE id = $1;
-- 
-- name: UpdateAward :exec
UPDATE awards SET name = $1, category = $2 WHERE id = $3;