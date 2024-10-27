-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY email;

-- name: CreateUser :one
INSERT INTO users (
    email, password, created_at, updated_at
) VALUES (
    $1, $2, NOW(), NOW()
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET 
    email = $2,
    updated_at = NOW()
WHERE email = $1
RETURNING *;