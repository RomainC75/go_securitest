package request_dto

type SignupRequest struct {
	LoginRequest
	// Birthday *time.Time `json:"birthday" `
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
