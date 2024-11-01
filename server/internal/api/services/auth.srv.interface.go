package services

import (
	db "server/db/sqlc"
	dto_req "server/internal/api/dtos/requests"
)

type IAuthSrv interface {
	Signup(signupData dto_req.UserCredsDto) (db.User, error)
	Login(loginData dto_req.UserCredsDto) (db.User, string, error)
}
