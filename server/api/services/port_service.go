package services

import (
	"server/api/repositories"
)

type PortSrv struct {
	portRepo repositories.PortRepositoryInterface
}

func NewPortSrv() *PortSrv {
	return &PortSrv{
		portRepo: repositories.NewPortRepo(),
	}
}

// func (userSrv *AuthSrv) CreatePortsSrv(ctx context.Context, address string, ports work_dto.PortTestScenario) (db.User, error) {
// 	foundUser, err := userSrv.userRepo.FindUserByEmail(ctx, user.Email)
// 	fmt.Println("==> found user : ", foundUser, err)
// 	if err == nil {
// 		return db.User{}, errors.New("email already used")
// 	}

// 	b, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return db.User{}, errors.New("error trying to encrypt the password")
// 	}

// 	userParams := db.CreateUserParams{
// 		Email:    user.Email,
// 		Password: string(b),
// 	}

// 	createdUser, err := userSrv.userRepo.CreateUser(ctx, userParams)
// 	utils.PrettyDisplay("createdUser", createdUser)
// 	if err != nil {
// 		return db.User{}, err
// 	}
// 	return createdUser, nil
// }

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
