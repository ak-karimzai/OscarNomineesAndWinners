-- name: CreateActor :one
INSERT INTO actors("name", "birth_year", "nationality") VALUES ($1, $2, $3) RETURNING id;

-- name: GetActor :one
SELECT * FROM actors WHERE id = $1 LIMIT 1;
 
-- name: ListActors :many
SELECT * FROM actors LIMIT $1 OFFSET $2;

-- name: DeleteActor :exec
DELETE FROM actors WHERE id = $1;

-- name: UpdateActor :exec
UPDATE actors SET name = $1, birth_year = $2, nationality = $3 WHERE id = $4;