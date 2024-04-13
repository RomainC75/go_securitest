package services

import (
	"context"
	"server/api/dto/requests"
	response_dto "server/api/dto/response"
	db "shared/db/sqlc"
)

type UserServiceInterface interface {
	CreateUserSrv(ctx context.Context, user requests.SignupRequest) (db.User, error)
	LoginSrv(ctx context.Context, user requests.LoginRequest) (response_dto.LoginResponse, error)
}
