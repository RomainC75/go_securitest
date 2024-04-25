-- name: GetScan :one
SELECT * FROM scans WHERE id = $1;

-- name: ListScansByUser :many
SELECT * FROM scans WHERE user_id = $1 ORDER BY executed_at;

-- name: CreateScan :one
INSERT INTO scans (
    user_id,
    executed_at,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: DeleteScan :one
DELETE FROM scans
WHERE id = $1 AND user_id = $2
RETURNING *;
