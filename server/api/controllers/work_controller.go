package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/api/controllers/ctrl_utils"
	"server/api/dto/requests"
	"server/api/services"
	"server/worker"
)

type WorkCtrl struct {
	userSrv services.UserSrv
}

func NewWorkCtrl() *WorkCtrl {
	return &WorkCtrl{
		userSrv: *services.NewUserSrv(),
	}
}

func (workCtrl *WorkCtrl) HandleWorkTest(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var req requests.SignupRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		json.NewEncoder(w).Encode("error ! ")
		return
	}
	fmt.Println("==> /work", req)

	ctx := context.Background()
	// createdUser, err := workCtrl.userSrv.CreateUserSrv(ctx, req)
	distributor := worker.Get()
	err = distributor.DistributeTaskSendVerifyEmail(ctx, &worker.PayloadSendVerifyEmail{
		Username: req.Email,
	})
	if err != nil {
		fmt.Println("ERR : ", err.Error())
	}

	// fmt.Println("=> <", req)
	ctrl_utils.SendJsonResponse(w, http.StatusCreated, map[string]any{"message": "created"})
}
