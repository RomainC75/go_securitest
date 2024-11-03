package repositories

import (
	db "server/db/sqlc"
	shared_dto "shared/dto"
)

type IScanRepo interface {
	CreateScan(userId int64, scenario int, scanData shared_dto.FullPortTestScenarioReq) (db.Scan, error)
}
