-- name: GetAddress :one
SELECT * FROM addresses WHERE id = $1;

-- name: GetAddressByEmail :one
SELECT * FROM addresses WHERE ip_addr = $1 LIMIT 1;

-- name: CreateAddress :one
INSERT INTO addresses (
    ip_addr,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: DeleteAddress :one
DELETE FROM addresses
WHERE id = $1
RETURNING *;