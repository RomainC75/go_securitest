package db

import (
	"context"
)

type Querier interface {
	CreateScan(ctx context.Context, arg CreateScanParams) (CreateScanRow, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, email string) error
	GetScan(ctx context.Context, userID int64) (GetScanRow, error)
	GetUser(ctx context.Context, email string) (User, error)
	ListScansByUser(ctx context.Context, userID int64) ([]ListScansByUserRow, error)
	ListUsers(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
