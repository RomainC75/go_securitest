package repositories

import (
	db "server/db/sqlc"
	dto_req "server/internal/api/dtos/requests"
)

type IAuthRepo interface {
	CreateUser(signupData dto_req.UserSignupDto) (db.User, error)
}
