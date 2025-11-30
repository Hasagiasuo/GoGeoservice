package repositories

import domain "geoserv/internal/domain/models"

type DriverRepository interface {
	baseRepository[domain.Driver]
	Add(name string) error
	GetByName(name string) (*domain.Driver, error)
	SetBusy(name string, is_busy bool) error
}
