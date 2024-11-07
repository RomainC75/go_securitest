package repositories

import (
	db "server/db/sqlc"
	shared_dto "shared/dto"
)

type IScanRepo interface {
	CreateScan(userId int64, scenario int, scanData shared_dto.FullPortTestScenarioReq) (db.Scan, error)
	GetScan(userId int64) (db.Scan, error)
	ListScansByUser(userId int64) ([]db.ListScansByUserRow, error)
}
