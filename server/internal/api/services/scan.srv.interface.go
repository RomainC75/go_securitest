package services

import (
	db "server/db/sqlc"
	shared_dto "shared/dto"
)

type IScanSrv interface {
	CreateScan(userId int, scenario int, reqData shared_dto.FullPortTestScenarioReq) error
	GetScan(userId int64) (db.Scan, error)
}
