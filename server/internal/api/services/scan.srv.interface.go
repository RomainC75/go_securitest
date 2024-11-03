package services

import shared_dto "shared/dto"

type IScanSrv interface {
	CreateScan(userId int, scenario int, reqData shared_dto.FullPortTestScenarioReq) error
}
