package services

import (
	"encoding/json"
	"server/internal/queue"
	shared_dto "shared/dto"
)

type ScanSrv struct {
	q *queue.SQueue
}

func NewScanSrv() *ScanSrv {
	return &ScanSrv{
		q: queue.GetQueue(),
	}
}

func (ss *ScanSrv) HandleScan(scenario int, reqData shared_dto.FullPortTestScenario) error {
	b, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	ss.q.Strategy.Push("azerty", b)
	// switch scenario {
	// case 1:

	// }

	return nil
}
