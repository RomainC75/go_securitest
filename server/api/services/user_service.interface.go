package services

import (
	"context"
	"server/api/dto/requests"
	db "server/db/sqlc"
)

type UserServiceInterface interface {
	CreateUserSrv(ctx context.Context, user requests.SignupRequest) (db.User, error)
}
