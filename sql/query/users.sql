-- name: CreateUser :one
INSERT INTO users(name)
    VALUES (?)
    RETURNING *;

-- name: Users :many
SELECT * FROM users;

-- name: UserById :one
SELECT *
    FROM users
    WHERE id = ?;

-- name: UserByName :one
SELECT *
    FROM users
    WHERE name = ?;

-- name: UserByApiKey :one
SELECT *
    FROM users
    WHERE api_key = ?;

-- name: UpdateUser :one
UPDATE users
    SET name = ?, api_key = ?
    WHERE id = ?
    RETURNING *;
