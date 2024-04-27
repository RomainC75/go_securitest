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

func (portRepo *PortRepository) CreatePorts(ctx context.Context, ports []db.CreatePortParams) ([]db.Port, error) {
	createdPorts := []db.Port{}
	for _, port := range ports {
		createdPort, err := portRepo.CreatePort(ctx, port)
		if err != nil {
			return createdPorts, err
		}
		createdPorts = append(createdPorts, createdPort)
	}
	return createdPorts, nil
}

func (portRepo *PortRepository) CreatePort(ctx context.Context, arg db.CreatePortParams) (db.Port, error) {
	port, err := (*portRepo.Store).CreatePort(ctx, arg)
	if err != nil {
		return db.Port{}, err
	}
	return port, nil
}

// func (userRepo *PortRepository) FindUserByEmail(ctx context.Context, email string) (db.User, error) {
// 	foundUser, err := (*userRepo.Store).GetUserByEmail(ctx, email)
// 	if err != nil {
// 		return db.User{}, err
// 	}
// 	return foundUser, nil
// }
