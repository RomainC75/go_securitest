package controllers

import (
	"context"
	"net/http"
	"server/api/controllers/ctrl_utils"
	request_dto "server/api/dto/request"
	response_dto "server/api/dto/response"
	"server/api/services"
)

type AuthCtrl struct {
	authSrv services.AuthSrv
}

func NewAuthCtrl() *AuthCtrl {
	return &AuthCtrl{
		authSrv: *services.NewUserSrv(),
	}
}

func (authCtrl *AuthCtrl) HandleSignupUser(w http.ResponseWriter, r *http.Request) {
	var reqBody request_dto.SignupRequest
	if err, status := ctrl_utils.CustomBodyValidator(w, r, &reqBody); err != nil {
		ctrl_utils.SendErrorResponse(w, status, err.Error(), ctrl_utils.ValidationErrorType)
		return
	}

	ctx := context.Background()
	createdUser, err := authCtrl.authSrv.CreateUserSrv(ctx, reqBody)
	if err != nil {
		ctrl_utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	signupResponse := response_dto.UserToSignupResponse(createdUser)
	ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"created": signupResponse})
}

func (authCtrl *AuthCtrl) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	var reqBody request_dto.LoginRequest
	if err, status := ctrl_utils.CustomBodyValidator(w, r, &reqBody); err != nil {
		ctrl_utils.SendErrorResponse(w, status, err.Error(), ctrl_utils.ValidationErrorType)
		return
	}

	ctx := context.Background()
	userResponse, err := authCtrl.authSrv.LoginSrv(ctx, reqBody)
	if err != nil {
		ctrl_utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"authorized": userResponse})
}
