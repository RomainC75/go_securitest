package db

import (
	"context"
	"database/sql"
)

const createIpRanges = `-- name: CreateIpRanges :one
INSERT INTO ip_ranges (scan_id, ip_min, ip_max)
VALUES ($1, $2, $3)
RETURNING id, scan_id, ip_min, ip_max
`

type CreateIpRangesParams struct {
	ScanID int64          `json:"scanId"`
	IpMin  string         `json:"ipMin"`
	IpMax  sql.NullString `json:"ipMax"`
}

func (q *Queries) CreateIpRanges(ctx context.Context, arg CreateIpRangesParams) (IpRange, error) {
	row := q.db.QueryRowContext(ctx, createIpRanges, arg.ScanID, arg.IpMin, arg.IpMax)
	var i IpRange
	err := row.Scan(
		&i.ID,
		&i.ScanID,
		&i.IpMin,
		&i.IpMax,
	)
	return i, err
}
