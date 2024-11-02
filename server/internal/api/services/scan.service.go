package services

import (
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

func (s *ScanSrv) HandleScan(scenario int, reqData shared_dto.FullPortTestScenario) error {

	switch scenario {
	case 1:

	}

	return nil
}
