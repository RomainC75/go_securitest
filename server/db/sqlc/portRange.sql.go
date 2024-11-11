package db

import (
	"context"
	"database/sql"
)

const createPortRanges = `-- name: CreatePortRanges :one
INSERT INTO port_ranges (scan_id, range_min, range_max)
VALUES ($1, $2, $3)
RETURNING id, scan_id, range_min, range_max, is_unique
`

type CreatePortRangesParams struct {
	ScanID   int64         `json:"scanId"`
	RangeMin int32         `json:"rangeMin"`
	RangeMax sql.NullInt32 `json:"rangeMax"`
}

func (q *Queries) CreatePortRanges(ctx context.Context, arg CreatePortRangesParams) (PortRange, error) {
	row := q.db.QueryRowContext(ctx, createPortRanges, arg.ScanID, arg.RangeMin, arg.RangeMax)
	var i PortRange
	err := row.Scan(
		&i.ID,
		&i.ScanID,
		&i.RangeMin,
		&i.RangeMax,
		&i.IsUnique,
	)
	return i, err
}
