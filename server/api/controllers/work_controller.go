package controllers

import (
	"fmt"
	"net/http"
	"server/api/controllers/ctrl_utils"
	request_dto "server/api/dto/request"
	"server/api/services"
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
		var portTestScenario request_dto.PortTestScenario
		if err, status := ctrl_utils.CustomBodyValidator(w, r, &portTestScenario); err != nil {
			ctrl_utils.SendErrorResponse(w, status, err.Error(), ctrl_utils.ValidationErrorType)
			return
		}
		fmt.Println("==> ", portTestScenario)
	default:
		ctrl_utils.SendErrorResponse(w, http.StatusBadRequest, "wrong scenario type", ctrl_utils.ValidationErrorType)
		return
	}

	// ctx := context.Background()
	// // createdUser, err := workCtrl.userSrv.CreateUserSrv(ctx, req)
	// distributor := events.Get()
	// opts := []asynq.Option{
	// 	asynq.MaxRetry(10),
	// 	asynq.ProcessIn(10 * time.Second),
	// 	// send to other queue :-)
	// 	asynq.Queue(events.CriticalQueue),
	// }
	// err = distributor.DistributeTaskSendVerifyEmail(ctx, &events.PayloadSendVerifyEmail{
	// 	Username: req.Email,
	// }, opts...)
	// if err != nil {
	// 	ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})
	// }

	// fmt.Println("=> <", req)
	ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})
}
