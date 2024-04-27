package services

import (
	"context"
	request_dto "server/api/dto/request"
	response_dto "server/api/dto/response"
	db "server/db/sqlc"
)

type PortServiceInterface interface {
	CreateUserSrv(ctx context.Context, user request_dto.SignupRequest) (db.User, error)
	LoginSrv(ctx context.Context, user request_dto.LoginRequest) (response_dto.LoginResponse, error)
}
