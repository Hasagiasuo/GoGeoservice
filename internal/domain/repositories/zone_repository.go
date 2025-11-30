package repositories

import domain "geoserv/internal/domain/models"

type ZoneRepository interface {
	baseRepository[domain.Zone]
	Add(zone_day string) error
}
