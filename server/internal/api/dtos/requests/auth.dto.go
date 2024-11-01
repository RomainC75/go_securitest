package dto_req

type UserCredsDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type WhoAmIDto struct {
	Token string `json:"token"`
}
