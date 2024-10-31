package services

import (
	db "server/db/sqlc"
	dto_req "server/internal/api/dtos/requests"
)

type IAuthSrv interface {
	Signup(signupData dto_req.UserSignupDto) (db.User, error)
}
