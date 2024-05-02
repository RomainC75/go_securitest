package repositories

import (
	"context"
	db "server/db/sqlc"
)

type ScanRepositoryInterface interface {
	CreateScan(ctx context.Context, userId int32) (db.Scan, error)
	// FindUserByEmail(ctx context.Context, email string) (db.User, error)
}
