package repositories

import (
	"context"
	db "server/db/sqlc"
	shared_dto "shared/dto"
)

type IScanRepo interface {
	CreateScan(ctx context.Context, userId int64, scenario int, scanData shared_dto.FullPortTestScenarioReq) (db.CreateScanRow, error)
	GetScan(userId int64) (db.GetScanRow, error)
	ListScansByUser(userId int64) ([]db.ListScansByUserRow, error)
}
