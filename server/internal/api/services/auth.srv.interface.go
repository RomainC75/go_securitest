package services

import (
	db "server/db/sqlc"
	"server/internal/api/dtos"
)

type IAuthSrv interface {
	Signup(signupData dtos.UserSignupDto) (db.User, error)
}
