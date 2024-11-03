-- name: GetScan :one
SELECT * FROM scans
LEFT JOIN port_ranges ON scans.id = port_ranges.scan_id
LEFT JOIN ip_ranges ON scans.id = ip_ranges.scan_id
WHERE scans.id = $1 LIMIT 1;

-- name: ListScansByUser :many
SELECT * FROM scans
LEFT JOIN port_ranges ON scans.id = port_ranges.scan_id
LEFT JOIN ip_ranges ON scans.id = ip_ranges.scan_id
WHERE scans.user_id = $1
ORDER BY scans.updated_at;

-- name: CreateScan :one
WITH inserted_scan AS (
    INSERT INTO scans (user_id, scenario, created_at, updated_at)
    VALUES ($1, $2, NOW(), NOW())
    RETURNING id
),
inserted_port_ranges AS (
    INSERT INTO port_ranges (scan_id, range_min, range_max)
    VALUES (id, $3, $4)
    RETURNING scan_id
),
inserted_ip_ranges AS (
    INSERT INTO ip_ranges (scan_id, ip_min, ip_max, is_unique)
    VALUES (id, $5, $6, $7)
)
SELECT * FROM scans WHERE id = (SELECT id FROM inserted_scan);

-- -- name: DeleteUser :exec
-- DELETE FROM users
-- WHERE email = $1;

-- -- name: UpdateUser :one
-- UPDATE users
-- SET 
--     email = $2,
--     updated_at = NOW()
-- WHERE email = $1
-- RETURNING *;