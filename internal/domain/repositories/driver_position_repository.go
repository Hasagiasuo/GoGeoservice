package repositories

import (
	"context"
	"geoserv/internal/domain/models"
)

type DriverPositionRepository interface {
	AddDriver(ctx context.Context, driver_position models.DriverPosition) error
	GetNerdyDriver(ctx context.Context, lon, lat, radius float64) ([]models.DriverPosition, error)
}
