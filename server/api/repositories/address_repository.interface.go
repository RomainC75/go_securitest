package repositories

import (
	"context"
	db "server/db/sqlc"

	"github.com/sqlc-dev/pqtype"
)

type AddressRepositoryInterface interface {
	CreateAddress(ctx context.Context, ipAddr pqtype.Inet) (db.Address, error)
	// FindUserByEmail(ctx context.Context, email string) (db.User, error)
}
