package services

import (
	db "server/db/sqlc"
	dto_req "server/internal/api/dtos/requests"
	"server/internal/api/repositories"
	"server/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthSrv struct {
	UserRepo repositories.IAuthRepo
}

func NewAuthSrv() *AuthSrv {
	return &AuthSrv{
		UserRepo: repositories.NewAuthRepo(),
	}
}

const BCRYPT_COST = 14

func (authSrv *AuthSrv) Signup(signupData dto_req.UserSignupDto) (db.User, string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), BCRYPT_COST)
	if err != nil {
		return db.User{}, "", err
	}
	signupData.Password = string(hashedPassword)

	createdUser, err := authSrv.UserRepo.CreateUser(signupData)
	if err != nil {
		return db.User{}, "", err
	}

	tokenString, err := utils.GenerateToken(createdUser)
	return createdUser, tokenString, err
}
