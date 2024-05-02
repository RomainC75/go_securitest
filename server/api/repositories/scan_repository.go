package repositories

import (
	"context"
	db "server/db/sqlc"
	"time"
)

type ScanRepository struct {
	Store *db.Store
}

func NewScanRepo() *ScanRepository {
	return &ScanRepository{
		Store: db.GetConnection(),
	}
}

func (scanRepo *ScanRepository) CreateScan(ctx context.Context, userId int32) (db.Scan, error) {
	now := time.Now()
	createdScan, err := (*scanRepo.Store).CreateScan(ctx, db.CreateScanParams{
		UserID:     userId,
		ExecutedAt: now,
		CreatedAt:  now,
		UpdatedAt:  now,
	})
	if err != nil {
		return db.Scan{}, err
	}
	return createdScan, nil
}
