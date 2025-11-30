package usecase

import (
	"geoserv/internal/domain/models"
	"geoserv/internal/domain/repositories"
)

type DriverUsecase struct {
	dr  repositories.DriverRepository
	dpr repositories.DriverPositionRepository
}

func NewDriverUsecase(dr repositories.DriverRepository, dpr repositories.DriverPositionRepository) *DriverUsecase {
	return &DriverUsecase{dr, dpr}
}

func (du *DriverUsecase) Add(name string) error {
	return du.dr.Add(name)
}

func (du *DriverUsecase) Get(id int) (models.Driver, error) {
	return du.dr.Get(id)
}

func (du *DriverUsecase) GetAll() ([]models.Driver, error) {
	return du.dr.GetAll()
}
