package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"server/api/controllers/ctrl_utils"
	"server/api/dto/requests"
	response_dto "server/api/dto/response"
	"server/api/services"
)

type AuthCtrl struct {
	userSrv services.UserSrv
}

func NewAuthCtrl() *AuthCtrl {
	return &AuthCtrl{
		userSrv: *services.NewUserSrv(),
	}
}

func (authCtrl *AuthCtrl) HandleSignupUser(w http.ResponseWriter, r *http.Request) {

	var req requests.SignupRequest

	if err, status := ctrl_utils.CustomValidator(w, r, &req); err != nil {
		ctrl_utils.SendErrorResponse(w, status, err.Error(), ctrl_utils.ValidationErrorType)
		return
	}

	ctx := context.Background()
	createdUser, err := authCtrl.userSrv.CreateUserSrv(ctx, req)
	if err != nil {
		ctrl_utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	signupResponse := response_dto.UserToSignupResponse(createdUser)

	ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"created": signupResponse})
}

func (authCtrl *AuthCtrl) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var req requests.SignupRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		ctrl_utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// ctx := context.Background()

}
