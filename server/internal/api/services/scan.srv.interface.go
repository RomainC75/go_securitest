package services

import (
	"context"
	db "server/db/sqlc"
	tx_types "server/db/sqlc/txTypes"
	shared_dto "shared/dto"
)

type IScanSrv interface {
	CreateScan(c context.Context, userId int, scenario int, reqData shared_dto.FullPortTestScenarioReq) (tx_types.CreatedScanTx, error)
	GetScan(userId int64) (db.GetScanRow, error)
}
