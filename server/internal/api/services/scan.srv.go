package services

import (
	"context"
	"encoding/json"
	db "server/db/sqlc"
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

func (ss *ScanSrv) ListScansByUser(userId int64) ([]db.ListScansByUserRow, error) {
	return ss.scanRepo.ListScansByUser(userId)
}

func (ss *ScanSrv) GetScan(userId int64) (db.GetScanRow, error) {
	return ss.scanRepo.GetScan(userId)
}

func (ss *ScanSrv) CreateScan(c context.Context, userId int, scenario int, reqData shared_dto.FullPortTestScenarioReq) error {

	createdScan, err := ss.scanRepo.CreateScan(c, int64(userId), scenario, reqData)
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
