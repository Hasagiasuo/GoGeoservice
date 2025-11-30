package usecase

import (
	"context"
	"geoserv/internal/domain/models"
	"geoserv/internal/domain/repositories"
)

type DriverPositionUsecase struct {
	dpr repositories.DriverPositionRepository
}

func NewDriverPositionUsecase(dpr repositories.DriverPositionRepository) *DriverPositionUsecase {
	return &DriverPositionUsecase{
		dpr,
	}
}

func (dpu *DriverPositionUsecase) AddDriverPosition(ctx context.Context, driver_position models.DriverPosition) error {
	return dpu.dpr.AddDriver(ctx, driver_position)
}

func (dpu *DriverPositionUsecase) GetDriversByRadius(ctx context.Context, lon, lat, radius float64) ([]models.DriverPosition, error) {
	return dpu.dpr.GetNerdyDriver(ctx, lon, lat, radius)
}
