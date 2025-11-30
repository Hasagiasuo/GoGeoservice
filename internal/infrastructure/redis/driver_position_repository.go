package redis

import (
	"context"
	"geoserv/internal/domain/models"
	"log"
	"strconv"

	red "github.com/redis/go-redis/v9"
)

type RedisDriverPositionRepository struct {
	client *RedisClient
	key    string
}

func NewRedisDriverPositionRepository(client *RedisClient) *RedisDriverPositionRepository {
	return &RedisDriverPositionRepository{
		client: client,
		key:    "drivers:positions",
	}
}

func (rr *RedisDriverPositionRepository) AddDriver(ctx context.Context, driver_position models.DriverPosition) error {
	_, err := rr.client.rdb.GeoAdd(ctx, rr.key, &red.GeoLocation{
		Name:      strconv.Itoa(driver_position.Id),
		Longitude: driver_position.Longitude,
		Latitude:  driver_position.Latitude,
	}).Result()
	log.Printf("Added driver position: %v", driver_position)
	return err
}

func (rr *RedisDriverPositionRepository) GetNerdyDriver(ctx context.Context, lon, lat, radius float64) ([]models.DriverPosition, error) {
	res, err := rr.client.rdb.GeoRadius(ctx, rr.key, lon, lat, &red.GeoRadiusQuery{
		Radius:    radius,
		Unit:      "km",
		WithCoord: true,
	}).Result()
	if err != nil {
		return nil, err
	}
	result := make([]models.DriverPosition, 0, len(res))
	for _, r := range res {
		id, err := strconv.Atoi(r.Name)
		if err != nil {
			log.Panicf("cannot parse id from name: %v", err)
			continue
		}
		result = append(result, models.DriverPosition{
			Id:        id,
			Latitude:  r.Latitude,
			Longitude: r.Longitude,
		})
	}
	return result, nil
}

