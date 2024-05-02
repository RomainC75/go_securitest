package services

import (
	"context"
	"server/api/repositories"
	db "server/db/sqlc"
	"shared/scenarios"
)

type ScanSrv struct {
	portRepo    repositories.PortRepositoryInterface
	addressRepo repositories.AddressRepositoryInterface
	scanRepo    repositories.ScanRepositoryInterface
}

func NewScanSrv() *ScanSrv {
	return &ScanSrv{
		portRepo:    repositories.NewPortRepo(),
		addressRepo: repositories.NewAddressRepo(),
		scanRepo:    repositories.NewScanRepo(),
	}
}

func (portSrv *ScanSrv) CreateScanSrv(ctx context.Context, scanResult scenarios.ScanResult) (db.Scan, error) {
	// scan

	// address

	// ports + address link
	// foundUser, err := portSrv.portRepo.CreatePort()
	// fmt.Println("==> found user : ", foundUser, err)
	// if err == nil {
	// 	return db.User{}, errors.New("email already used")
	// }

	// return db.Scan{}, nil

	// scan + address link
	return db.Scan{}, nil
}

// func (userSrv *AuthSrv) LoginSrv(ctx context.Context, user request_dto.LoginRequest) (response_dto.LoginResponse, error) {
// 	foundUser, err := userSrv.userRepo.FindUserByEmail(ctx, user.Email)
// 	if err != nil || foundUser.Email == "" {
// 		return response_dto.LoginResponse{}, err
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
// 	if err != nil {
// 		return response_dto.LoginResponse{}, err
// 	}

// 	token, err := encrypt.Generate(foundUser)
// 	if err != nil {
// 		return response_dto.LoginResponse{}, err
// 	}
// 	response := response_dto.LoginResponse{
// 		ID:    foundUser.ID,
// 		Email: user.Email,
// 		Token: token,
// 	}

// 	return response, nil
// }
