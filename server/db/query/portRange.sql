-- name: CreatePortRanges :one
INSERT INTO port_ranges (scan_id, range_min, range_max)
VALUES ($1, $2, $3)
RETURNING *;
