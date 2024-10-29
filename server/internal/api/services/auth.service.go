package services

import (
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

func (authSrv *AuthSrv) Signup(signupData dtos.UserSignupDto) (db.User, error) {
	return authSrv.UserRepo.CreateUser(signupData)
}
