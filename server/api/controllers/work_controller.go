package controllers

import (
	"context"
	"fmt"
	"net/http"
	"server/api/controllers/ctrl_utils"
	request_dto "server/api/dto/request"
	"server/api/services"
	"shared/events"

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
	workCode := r.PathValue("work_code")
	fmt.Println("=> WK ! ", workCode)

	switch workCode {
	case "1":
		var portTestScenario request_dto.FullPortTestScenario
		if err, status := ctrl_utils.CustomBodyValidator(w, r, &portTestScenario); err != nil {
			ctrl_utils.SendErrorResponse(w, status, err.Error(), ctrl_utils.ValidationErrorType)
			return
		}
		fmt.Println("==> ", portTestScenario)

		ctx := context.Background()
		distributor := events.Get()
		opts := []asynq.Option{
			asynq.MaxRetry(10),
			// asynq.ProcessIn(time.Second),
			// + send to other queue :-)
			asynq.Queue(events.CriticalQueue),
		}

		err := distributor.DistributeTaskSendWork(
			ctx,
			&portTestScenario.PortTestScenario,
			opts...,
		)
		if err != nil {
			ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})
		}

	default:
		ctrl_utils.SendErrorResponse(w, http.StatusBadRequest, "wrong scenario type", ctrl_utils.ValidationErrorType)
		return
	}

	// fmt.Println("=> <", req)
	ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})
}
