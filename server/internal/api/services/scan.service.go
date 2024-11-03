package services

import (
	"encoding/json"
	"server/internal/queue"
	shared_dto "shared/dto"
	"time"

	"github.com/google/uuid"
)

type ScanSrv struct {
	q *queue.SQueue
}

func NewScanSrv() *ScanSrv {
	return &ScanSrv{
		q: queue.GetQueue(),
	}
}

func (ss *ScanSrv) HandleScan(scenario int, reqData shared_dto.FullPortTestScenarioReq) error {
	eventReqData := shared_dto.Event{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		Content:   reqData,
	}

	b, err := json.Marshal(eventReqData)
	if err != nil {
		return err
	}

	ss.q.Strategy.Push("azerty", b)

	return nil
}
