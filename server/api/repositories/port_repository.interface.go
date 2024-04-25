package repositories

import (
	"context"
	db "server/db/sqlc"
)

type PortRepositoryInterface interface {
	CreatePort(ctx context.Context, arg db.CreatePortParams) (db.Port, error)
	// FindUserByEmail(ctx context.Context, email string) (db.User, error)
}
