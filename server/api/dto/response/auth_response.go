package response_dto

import (
	db "server/db/sqlc"
	"time"
)

type SignupResponse struct {
	ID        int32     `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserToSignupResponse(user db.User) SignupResponse {
	return SignupResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type LoginResponse struct {
	ID    int32  `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
