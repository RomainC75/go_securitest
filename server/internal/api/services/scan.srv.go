package services

import (
	"context"
	"encoding/json"
	db "server/db/sqlc"
	tx_types "server/db/sqlc/txTypes"
	"server/internal/api/repositories"
	"server/internal/queue"
	shared_dto "shared/dto"
	shared_utils "shared/utils"
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

func (ss *ScanSrv) CreateScan(c context.Context, userId int, scenario int, reqData shared_dto.FullPortTestScenarioReq) (tx_types.CreatedScanTx, error) {

	createdScan, err := ss.scanRepo.CreateTxScan(c, int64(userId), scenario, reqData)
	if err != nil {
		return tx_types.CreatedScanTx{}, err
	}
	shared_utils.PrettyDisplay("Created Scan", createdScan)
	eventReqData := shared_dto.Event{
		Id:        createdScan.Scan.ID,
		CreatedAt: createdScan.Scan.CreatedAt,
		Scenario:  scenario,
		Content:   reqData,
	}

	b, err := json.Marshal(eventReqData)
	if err != nil {
		return tx_types.CreatedScanTx{}, err
	}

	ss.q.Strategy.Push("azerty", b)

	return createdScan, nil
}
