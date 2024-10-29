package repositories

import (
	"context"
	db "server/db/sqlc"
	"server/internal/api/dtos"
)

type IAuthRepo interface {
	CreateUser(ctx context.Context, signupData dtos.UserSignupDto) (db.User, error)
}
