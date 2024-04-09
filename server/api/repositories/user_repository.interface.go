package repositories

import (
	"context"
	db "server/db/sqlc"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	FindUserByEmail(ctx context.Context, email string) (db.User, error)
}
