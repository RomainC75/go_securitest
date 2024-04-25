-- name: GetPort :one
SELECT * FROM ports WHERE id = $1;

-- name: GetPortsByIpId :many
SELECT * FROM ports WHERE ip_addr_id = $1 ORDER BY port;

-- name: CreatePort :one
INSERT INTO ports (
    ip_addr_id,
    port,
    state,
    executed_at
) VALUES (
    $1, $2, $3, $4
) RETURNING *;