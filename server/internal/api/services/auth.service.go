package services

import (
	"context"
	"fmt"
	db "server/db/sqlc"
	"server/internal/api/dtos"
	"server/internal/api/repositories"
)

type AuthSrv struct {
	UserRepo repositories.IAuthRepo
}

func NewAuthSrv() *AuthSrv {
	return &AuthSrv{
		UserRepo: repositories.NewAuthRepo(),
	}
}

func (authSrv *AuthSrv) Signup(ctx context.Context, signupData dtos.UserSignupDto) (db.User, error) {
	fmt.Println("SRV")
	return authSrv.UserRepo.CreateUser(ctx, signupData)
}
