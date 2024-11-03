package services

import (
	"encoding/json"
	"server/internal/api/repositories"
	"server/internal/queue"
	shared_dto "shared/dto"
)

type ScanSrv struct {
	q        *queue.SQueue
	scanRepo repositories.IScanRepo
}

func NewScanSrv() *ScanSrv {
	return &ScanSrv{
		q:        queue.GetQueue(),
		scanRepo: repositories.NewScanRepo(),
	}
}

func (ss *ScanSrv) CreateScan(userId int, scenario int, reqData shared_dto.FullPortTestScenarioReq) error {

	createdScan, err := ss.scanRepo.CreateScan(int64(userId), scenario, reqData)
	if err != nil {
		return err
	}
	eventReqData := shared_dto.Event{
		Id:        createdScan.ID,
		CreatedAt: createdScan.CreatedAt,
		Content:   reqData,
	}

	b, err := json.Marshal(eventReqData)
	if err != nil {
		return err
	}

	ss.q.Strategy.Push("azerty", b)

	return nil
}
