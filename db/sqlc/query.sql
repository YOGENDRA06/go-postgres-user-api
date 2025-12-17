-- name: CreateUser :one
INSERT INTO users (name, dob)
VALUES ($1, $2)
RETURNING id, name, dob, created_at;

-- name: GetUserByID :one
SELECT id, name, dob, created_at
FROM users
WHERE id = $1;

-- name: GetUsers :many
SELECT id, name, dob, created_at
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET name = $2,
    dob  = $3
WHERE id = $1
RETURNING id, name, dob, created_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
