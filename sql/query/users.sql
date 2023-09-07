-- name: CreateUser :one
INSERT INTO users(name)
VALUES (?)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;
