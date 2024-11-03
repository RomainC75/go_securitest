package db

import (
	"context"
	"database/sql"
	"time"
)

const createScan = `-- name: CreateScan :one
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
SELECT id, user_id, scenario, created_at, updated_at FROM scans WHERE id = (SELECT id FROM inserted_scan)
`

type CreateScanParams struct {
	UserID   int64          `json:"userId"`
	Scenario int32          `json:"scenario"`
	RangeMin int32          `json:"rangeMin"`
	RangeMax sql.NullInt32  `json:"rangeMax"`
	IpMin    string         `json:"ipMin"`
	IpMax    sql.NullString `json:"ipMax"`
	IsUnique sql.NullBool   `json:"isUnique"`
}

func (q *Queries) CreateScan(ctx context.Context, arg CreateScanParams) (Scan, error) {
	row := q.db.QueryRowContext(ctx, createScan,
		arg.UserID,
		arg.Scenario,
		arg.RangeMin,
		arg.RangeMax,
		arg.IpMin,
		arg.IpMax,
		arg.IsUnique,
	)
	var i Scan
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Scenario,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getScan = `-- name: GetScan :one
SELECT scans.id, user_id, scenario, created_at, updated_at, port_ranges.id, port_ranges.scan_id, range_min, range_max, port_ranges.is_unique, ip_ranges.id, ip_ranges.scan_id, ip_min, ip_max, ip_ranges.is_unique FROM scans
LEFT JOIN port_ranges ON scans.id = port_ranges.scan_id
LEFT JOIN ip_ranges ON scans.id = ip_ranges.scan_id
WHERE scans.id = $1 LIMIT 1
`

type GetScanRow struct {
	ID         int64          `json:"id"`
	UserID     int64          `json:"userId"`
	Scenario   int32          `json:"scenario"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	ID_2       sql.NullInt64  `json:"id2"`
	ScanID     sql.NullInt64  `json:"scanId"`
	RangeMin   sql.NullInt32  `json:"rangeMin"`
	RangeMax   sql.NullInt32  `json:"rangeMax"`
	IsUnique   sql.NullBool   `json:"isUnique"`
	ID_3       sql.NullInt64  `json:"id3"`
	ScanID_2   sql.NullInt64  `json:"scanId2"`
	IpMin      sql.NullString `json:"ipMin"`
	IpMax      sql.NullString `json:"ipMax"`
	IsUnique_2 sql.NullBool   `json:"isUnique2"`
}

func (q *Queries) GetScan(ctx context.Context, id int64) (GetScanRow, error) {
	row := q.db.QueryRowContext(ctx, getScan, id)
	var i GetScanRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Scenario,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID_2,
		&i.ScanID,
		&i.RangeMin,
		&i.RangeMax,
		&i.IsUnique,
		&i.ID_3,
		&i.ScanID_2,
		&i.IpMin,
		&i.IpMax,
		&i.IsUnique_2,
	)
	return i, err
}

const listScansByUser = `-- name: ListScansByUser :many
SELECT scans.id, user_id, scenario, created_at, updated_at, port_ranges.id, port_ranges.scan_id, range_min, range_max, port_ranges.is_unique, ip_ranges.id, ip_ranges.scan_id, ip_min, ip_max, ip_ranges.is_unique FROM scans
LEFT JOIN port_ranges ON scans.id = port_ranges.scan_id
LEFT JOIN ip_ranges ON scans.id = ip_ranges.scan_id
WHERE scans.user_id = $1
ORDER BY scans.updated_at
`

type ListScansByUserRow struct {
	ID         int64          `json:"id"`
	UserID     int64          `json:"userId"`
	Scenario   int32          `json:"scenario"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	ID_2       sql.NullInt64  `json:"id2"`
	ScanID     sql.NullInt64  `json:"scanId"`
	RangeMin   sql.NullInt32  `json:"rangeMin"`
	RangeMax   sql.NullInt32  `json:"rangeMax"`
	IsUnique   sql.NullBool   `json:"isUnique"`
	ID_3       sql.NullInt64  `json:"id3"`
	ScanID_2   sql.NullInt64  `json:"scanId2"`
	IpMin      sql.NullString `json:"ipMin"`
	IpMax      sql.NullString `json:"ipMax"`
	IsUnique_2 sql.NullBool   `json:"isUnique2"`
}

func (q *Queries) ListScansByUser(ctx context.Context, userID int64) ([]ListScansByUserRow, error) {
	rows, err := q.db.QueryContext(ctx, listScansByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListScansByUserRow{}
	for rows.Next() {
		var i ListScansByUserRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Scenario,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.ScanID,
			&i.RangeMin,
			&i.RangeMax,
			&i.IsUnique,
			&i.ID_3,
			&i.ScanID_2,
			&i.IpMin,
			&i.IpMax,
			&i.IsUnique_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
