package repositories

import (
	"context"
	db "server/db/sqlc"
	dto_req "server/internal/api/dtos/requests"
)

type UserRepo struct {
	Store *db.Store
}

func NewAuthRepo() *UserRepo {
	return &UserRepo{
		Store: db.GetConnection(),
	}
}

func (userRepo *UserRepo) CreateUser(signupData dto_req.UserCredsDto) (db.User, error) {
	ctx := context.Background()

	createdUser := db.CreateUserParams{
		Email:    signupData.Email,
		Password: signupData.Password,
	}
	return (*userRepo.Store).CreateUser(ctx, createdUser)
}

func (userRepo *UserRepo) GetUser(email string) (db.User, error) {
	ctx := context.Background()
	return (*userRepo.Store).GetUser(ctx, email)
}
