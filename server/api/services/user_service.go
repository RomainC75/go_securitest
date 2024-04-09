package services

import (
	"context"
	"errors"
	"fmt"
	"server/api/dto/requests"
	"server/api/repositories"
	db "server/db/sqlc"
	"server/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserSrv struct {
	userRepo repositories.UserRepositoryInterface
}

func NewUserSrv() *UserSrv {
	return &UserSrv{
		userRepo: repositories.NewUserRepo(),
	}
}

func (userSrv *UserSrv) CreateUserSrv(ctx context.Context, user requests.SignupRequest) (db.User, error) {
	foundUser, err := userSrv.userRepo.FindUserByEmail(ctx, user.Email)
	fmt.Println("==> found user : ", foundUser, err)
	if err == nil {
		return db.User{}, errors.New("email already used")
	}

	b, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.User{}, errors.New("error trying to encrypt the password")
	}

	userParams := db.CreateUserParams{
		Email:    user.Email,
		Password: string(b),
	}

	createdUser, err := userSrv.userRepo.CreateUser(ctx, userParams)
	utils.PrettyDisplay("createdUser", createdUser)
	if err != nil {
		return db.User{}, err
	}
	return createdUser, nil
}
