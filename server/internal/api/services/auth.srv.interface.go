package services

import (
	"context"
	db "server/db/sqlc"
	"server/internal/api/dtos"
)

type IAuthSrv interface {
	Signup(ctx context.Context, signupData dtos.UserSignupDto) (db.User, error)
}
