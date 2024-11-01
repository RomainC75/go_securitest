package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	dto_req "server/internal/api/dtos/requests"
	dto_res "server/internal/api/dtos/responses"
	"server/internal/api/services"
	"server/utils"

	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	AuthSrv services.IAuthSrv
}

func NewAuthController() *AuthController {
	return &AuthController{
		AuthSrv: services.NewAuthSrv(),
	}
}

var validate *validator.Validate

func (c *AuthController) HandleAuthSignup(w http.ResponseWriter, r *http.Request) {
	var u dto_req.UserCredsDto

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := c.AuthSrv.Signup(u)
	if err != nil {
		fmt.Println("3")
		utils.SendError(w, 401, err)
		return
	}

	// err := validate.Struct(mystruct)
	// validationErrors := err.(validator.ValidationErrors)
	utils.SendJson(w, 200, map[string]any{
		"message": "created",
		"user": dto_res.UserSignupDtoRes{
			Id:    createdUser.ID,
			Email: createdUser.Email,
		},
	})
}

func (c *AuthController) HandleAuthLogin(w http.ResponseWriter, r *http.Request) {
	var u dto_req.UserCredsDto

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundUser, token, err := c.AuthSrv.Login(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendJson(w, 200, map[string]any{
		"message": "created",
		"user": dto_res.UserSignupDtoRes{
			Id:    foundUser.ID,
			Email: foundUser.Email,
		},
		"token": token,
	})
}
