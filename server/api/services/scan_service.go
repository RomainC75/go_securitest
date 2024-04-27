package services

import (
	"context"
	"server/api/repositories"
	db "server/db/sqlc"
	work_dto "shared/dto"
)

type ScanSrv struct {
	portRepo    repositories.PortRepositoryInterface
	addressRepo repositories.AddressRepositoryInterface
}

func NewScanSrv() *ScanSrv {
	return &ScanSrv{
		portRepo:    repositories.NewPortRepo(),
		addressRepo: repositories.NewAddressRepo(),
	}
}

func (portSrv *ScanSrv) CreateScanSrv(ctx context.Context, address string, ports work_dto.PortTestScenario) (db.Scan, error) {
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
