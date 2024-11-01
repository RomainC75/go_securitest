package repositories

import (
	db "server/db/sqlc"
	dto_req "server/internal/api/dtos/requests"
)

type IAuthRepo interface {
	CreateUser(signupData dto_req.UserCredsDto) (db.User, error)
	GetUser(email string) (db.User, error)
}
