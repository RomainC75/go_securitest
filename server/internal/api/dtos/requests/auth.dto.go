package dto_req

type UserCredsDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type WhoAmIDto struct {
	Token string `json:"token"`
}
