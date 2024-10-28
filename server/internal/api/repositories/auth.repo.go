package repositories

import (
	"context"
	db "server/db/sqlc"
	"server/internal/api/dtos"
)

type UserRepo struct {
	Store *db.Store
}

func NewAuthRepo() *UserRepo {
	return &UserRepo{
		Store: db.GetConnection(),
	}
}

func (userRepo *UserRepo) CreateUser(signupData dtos.UserSignupDto) (db.User, error) {
	ctx := context.Background()
	createdUser := db.CreateUserParams{
		Email:    signupData.Email,
		Password: signupData.Password,
	}
	return (*userRepo.Store).CreateUser(ctx, createdUser)
}
