package services

import (
	"context"
	db "server/db/sqlc"
	shared_dto "shared/dto"
)

type IScanSrv interface {
	CreateScan(c context.Context, userId int, scenario int, reqData shared_dto.FullPortTestScenarioReq) error
	GetScan(userId int64) (db.GetScanRow, error)
}
