package repositories

import (
	db "server/db/sqlc"
	"server/internal/api/dtos"
)

type IAuthRepo interface {
	CreateUser(signupData dtos.UserSignupDto) (db.User, error)
}
