-- name: GetUser :one
SELECT * FROM users WHERE username = $1;