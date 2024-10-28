package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/api/dtos"
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

	var u dtos.UserSignupDto

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("-->", u)

	// err := validate.Struct(mystruct)
	// validationErrors := err.(validator.ValidationErrors)
	utils.SendJson(w, 200, map[string]any{
		"message": u,
	})
}

func HandleAuthLogin(w http.ResponseWriter, r *http.Request) {

}
