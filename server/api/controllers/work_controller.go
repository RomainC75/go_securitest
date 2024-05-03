package controllers

import (
	"context"
	"net/http"
	"server/api/controllers/ctrl_utils"
	"server/api/services"
	work_dto "shared/dto"
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
	userId := r.Context().Value("user_id").(int32)

	switch workCode {
	case "1":
		var portTestScenario work_dto.FullPortTestScenario
		if err, status := ctrl_utils.CustomBodyValidator(w, r, &portTestScenario); err != nil {
			ctrl_utils.SendErrorResponse(w, status, err.Error(), ctrl_utils.ValidationErrorType)
			return
		}

		ctx := context.Background()
		distributor := events.Get()
		opts := []asynq.Option{
			asynq.MaxRetry(10),
			// asynq.ProcessIn(time.Second),
			// + send to other queue :-)
			asynq.Queue(string(events.CriticalQueueReq)),
		}

		scenarioRequest := work_dto.PortTestScenarioRequest{
			UserId:    userId,
			IPRange:   portTestScenario.PortTestScenario.IPRange,
			PortRange: portTestScenario.PortTestScenario.PortRange,
		}

		err := distributor.DistributeTaskSendWork(
			ctx,
			&scenarioRequest,
			events.PortScanner,
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
