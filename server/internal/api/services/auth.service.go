package services

import (
	"errors"
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

func (authSrv *AuthSrv) Signup(signupData dto_req.UserCredsDto) (db.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), BCRYPT_COST)
	if err != nil {
		return db.User{}, err
	}
	signupData.Password = string(hashedPassword)
	return authSrv.UserRepo.CreateUser(signupData)
}

func (authSrv *AuthSrv) Login(loginData dto_req.UserCredsDto) (db.User, string, error) {
	foundUser, err := authSrv.UserRepo.GetUser(loginData.Email)
	if err != nil {
		return db.User{}, "", errors.New("invalid email / password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginData.Password))
	if err != nil {
		return db.User{}, "", errors.New("invalid email / password")
	}

	tokenString, err := utils.GenerateToken(foundUser)
	if err != nil {
		return db.User{}, "", err
	}

	return foundUser, tokenString, nil
}
