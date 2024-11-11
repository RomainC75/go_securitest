-- name: CreateIpRanges :one
INSERT INTO ip_ranges (scan_id, ip_min, ip_max)
VALUES ($1, $2, $3)
RETURNING *;