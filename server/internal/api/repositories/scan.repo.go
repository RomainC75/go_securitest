package repositories

import (
	"context"
	"database/sql"
	db "server/db/sqlc"
	tx_types "server/db/sqlc/txTypes"
	shared_dto "shared/dto"
	shared_utils "shared/utils"
)

type ScanRepo struct {
	Store *db.Store
}

func NewScanRepo() *ScanRepo {
	return &ScanRepo{
		Store: db.GetConnection(),
	}
}

func (scanRepo *ScanRepo) CreateScan(ctx context.Context, userId int64, scenario int, scanData shared_dto.FullPortTestScenarioReq) (db.CreateScanRow, error) {
	scanToCreate := db.CreateScanParams{
		UserID:   userId,
		Scenario: int32(scenario),
		RangeMin: int32(scanData.PortTestScenario.PortRange.Min),
		IpMin:    scanData.PortTestScenario.IPRange.IpMin,
	}
	if scanData.PortTestScenario.PortRange.Max != nil {
		scanToCreate.RangeMax = sql.NullInt32{Int32: *scanData.PortTestScenario.PortRange.Max, Valid: true}
	}
	if scanData.PortTestScenario.IPRange.IpMax != nil {
		scanToCreate.IpMax = sql.NullString{String: *scanData.PortTestScenario.IPRange.IpMax, Valid: true}
	}
	shared_utils.PrettyDisplay("db.CreateScanParams : ", scanToCreate)
	res, err := (*scanRepo.Store).CreateScan(ctx, scanToCreate)
	shared_utils.PrettyDisplay("res : ", res)
	return res, err
}

func (scanRepo *ScanRepo) CreateTxScan(ctx context.Context, userId int64, scenario int, scanData shared_dto.FullPortTestScenarioReq) (tx_types.CreatedScanTx, error) {
	var result tx_types.CreatedScanTx
	err := (*scanRepo.Store).ExecTx(ctx, func(q *db.Queries) error {
		var err error
		createdScan, err := q.CreateScanTx(ctx, db.CreateScanTxParams{
			UserID:   userId,
			Scenario: int32(scenario),
		})
		if err != nil {
			return err
		}

		ipRangesParams := db.CreateIpRangesParams{
			ScanID: createdScan.ID,
			IpMin:  scanData.PortTestScenario.IPRange.IpMin,
		}
		if scanData.PortTestScenario.IPRange.IpMax != nil {
			ipRangesParams.IpMax = sql.NullString{String: *scanData.PortTestScenario.IPRange.IpMax, Valid: true}
		}
		createdIpRange, err := q.CreateIpRanges(ctx, ipRangesParams)
		if err != nil {
			return err
		}

		portRangesParams := db.CreatePortRangesParams{
			ScanID:   createdScan.ID,
			RangeMin: scanData.PortTestScenario.PortRange.Min,
		}
		if scanData.PortTestScenario.PortRange.Max != nil {
			portRangesParams.RangeMax = sql.NullInt32{Int32: *scanData.PortTestScenario.PortRange.Max, Valid: true}
		}
		createdPortRange, err := q.CreatePortRanges(ctx, portRangesParams)
		if err != nil {
			return err
		}
		result.Scan = createdScan
		result.IpRange = createdIpRange
		result.PortRange = createdPortRange
		return nil
	})
	return result, err
}

func (scanRepo *ScanRepo) GetScan(userId int64) (db.GetScanRow, error) {
	ctx := context.Background()
	return (*scanRepo.Store).GetScan(ctx, userId)
}

func (scanRepo *ScanRepo) ListScansByUser(userId int64) ([]db.ListScansByUserRow, error) {
	ctx := context.Background()
	return (*scanRepo.Store).ListScansByUser(ctx, userId)
}
