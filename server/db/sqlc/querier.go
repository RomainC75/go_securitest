package db

import (
	"context"
)

type Querier interface {
	CreateScan(ctx context.Context, arg CreateScanParams) (Scan, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, email string) error
	// LEFT JOIN port_ranges ON scans.id = port_ranges.scan_id
	// LEFT JOIN ip_ranges ON scans.id = ip_ranges.scan_id
	GetScan(ctx context.Context, userID int64) (Scan, error)
	GetUser(ctx context.Context, email string) (User, error)
	ListScansByUser(ctx context.Context, userID int64) ([]ListScansByUserRow, error)
	ListUsers(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
