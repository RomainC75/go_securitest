package repositories

import (
	"context"
	db "server/db/sqlc"
)

type PortRepository struct {
	Store *db.Store
}

func NewPortRepo() *PortRepository {
	return &PortRepository{
		Store: db.GetConnection(),
	}
}

func (userRepo *PortRepository) CreatePort(ctx context.Context, arg db.CreatePortParams) (db.Port, error) {
	user, err := (*userRepo.Store).CreatePort(ctx, arg)
	if err != nil {
		return db.Port{}, err
	}
	return user, nil
}

// func (userRepo *PortRepository) FindUserByEmail(ctx context.Context, email string) (db.User, error) {
// 	foundUser, err := (*userRepo.Store).GetUserByEmail(ctx, email)
// 	if err != nil {
// 		return db.User{}, err
// 	}
// 	return foundUser, nil
// }
