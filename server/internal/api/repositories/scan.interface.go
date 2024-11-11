package repositories

import (
	"context"
	db "server/db/sqlc"
	tx_types "server/db/sqlc/txTypes"
	shared_dto "shared/dto"
)

type IScanRepo interface {
	CreateScan(ctx context.Context, userId int64, scenario int, scanData shared_dto.FullPortTestScenarioReq) (db.CreateScanRow, error)
	CreateTxScan(ctx context.Context, userId int64, scenario int, scanData shared_dto.FullPortTestScenarioReq) (tx_types.CreatedScanTx, error)
	GetScan(userId int64) (db.GetScanRow, error)
	ListScansByUser(userId int64) ([]db.ListScansByUserRow, error)
}
