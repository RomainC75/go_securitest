package controllers

import (
	"encoding/json"
	"net/http"
	dto_req "server/internal/api/dtos/requests"
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
	var u dto_req.UserSignupDto

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := c.AuthSrv.Signup(u)
	if err != nil {
		utils.SendError(w, 401, err)
		return
	}

	// err := validate.Struct(mystruct)
	// validationErrors := err.(validator.ValidationErrors)
	utils.SendJson(w, 200, map[string]any{
		"message": createdUser,
	})
}

func HandleAuthLogin(w http.ResponseWriter, r *http.Request) {

}
