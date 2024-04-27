package repositories

import (
	"context"
	db "server/db/sqlc"
	"time"

	"github.com/sqlc-dev/pqtype"
)

type AddressRepository struct {
	Store *db.Store
}

func NewAddressRepo() *AddressRepository {
	return &AddressRepository{
		Store: db.GetConnection(),
	}
}

func (addressRepo *AddressRepository) CreateAddress(ctx context.Context, ipAddr pqtype.Inet) (db.Address, error) {
	now := time.Now()
	createdAddress, err := (*addressRepo.Store).CreateAddress(ctx, db.CreateAddressParams{
		IpAddr:    ipAddr,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return db.Address{}, err
	}
	return createdAddress, nil
}

// func (addressRepo *PortRepository) FindUserByEmail(ctx context.Context, email string) (db.User, error) {
// 	foundUser, err := (*addressRepo.Store).GetUserByEmail(ctx, email)
// 	if err != nil {
// 		return db.User{}, err
// 	}
// 	return foundUser, nil
// }
