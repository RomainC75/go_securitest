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
	"time"

	"github.com/hibiken/asynq"
)

type WorkCtrl struct {
	userSrv services.AuthSrv
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
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		// send to other queue :-)
		asynq.Queue(worker.CriticalQueue),
	}
	err = distributor.DistributeTaskSendVerifyEmail(ctx, &worker.PayloadSendVerifyEmail{
		Username: req.Email,
	}, opts...)
	if err != nil {
		ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})
	}

	// fmt.Println("=> <", req)
	ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})
}
