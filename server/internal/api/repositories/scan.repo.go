package repositories

import (
	"context"
	"database/sql"
	db "server/db/sqlc"
	shared_dto "shared/dto"
)

type ScanRepo struct {
	Store *db.Store
}

func NewScanRepo() *ScanRepo {
	return &ScanRepo{
		Store: db.GetConnection(),
	}
}

func (scanRepo *ScanRepo) CreateScan(userId int64, scenario int, scanData shared_dto.FullPortTestScenarioReq) (db.Scan, error) {
	ctx := context.Background()

	scanToCreate := db.CreateScanParams{
		UserID:   userId,
		Scenario: int32(scenario),
		RangeMin: int32(scanData.PortTestScenario.PortRange.Min),
		RangeMax: sql.NullInt32{
			Int32: scanData.PortTestScenario.PortRange.Max.Int32,
			Valid: scanData.PortTestScenario.PortRange.Max.Valid,
		},
		IpMin: scanData.PortTestScenario.IPRange.IpMin,
		IpMax: sql.NullString{
			String: scanData.PortTestScenario.IPRange.IpMax.String,
			Valid:  scanData.PortTestScenario.IPRange.IpMax.Valid,
		},
	}
	return (*scanRepo.Store).CreateScan(ctx, scanToCreate)
}

func (scanRepo *ScanRepo) GetUser(email string) (db.User, error) {
	ctx := context.Background()
	return (*scanRepo.Store).GetUser(ctx, email)
}
