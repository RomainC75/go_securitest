package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	validator_helper "server/internal/api/dtos/validator"
	"server/internal/api/services"
	"server/internal/queue"
	"server/utils"
	shared_dto "shared/dto"
	shared_utils "shared/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type AnalyseCtrl struct {
	Queue   *queue.SQueue
	v       *validator.Validate
	scanSrv *services.ScanSrv
}

func NewAnalyseCtrl() *AnalyseCtrl {
	return &AnalyseCtrl{
		Queue:   queue.GetQueue(),
		v:       validator_helper.GetValidate(),
		scanSrv: services.NewScanSrv(),
	}
}

func (c *AnalyseCtrl) HandleAnalyse(w http.ResponseWriter, r *http.Request) {
	workCode := r.PathValue("scenario")
	scenarioNum, err := strconv.Atoi(workCode)

	if err != nil {
		http.Error(w, "scan scenario should be a number", http.StatusBadRequest)
		return
	}
	fmt.Println("=> WK ! ", scenarioNum)

	var u shared_dto.FullPortTestScenario

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.v.Struct(u)
	if err != nil {
		logrus.Warnf("validator error : %s \n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shared_utils.PrettyDisplay("body ", u)

	err = c.scanSrv.HandleScan(scenarioNum, u)
	if err != nil {
		logrus.Warnf("err : %s \n", err.Error())
	}

	utils.SendJson(w, http.StatusOK, map[string]any{
		"message": "/analyse",
	})
}
