package redis

import (
	"context"
	"fmt"
	"geoserv/internal/domain/models"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisDriverRepository struct {
	client *RedisClient
}

func NewRedisDriverRepository(client *RedisClient) *RedisDriverRepository {
	return &RedisDriverRepository{
		client,
	}
}

func (rdr *RedisDriverRepository) SetDriverState(ctx context.Context, driver_id int, state models.DriverState) error {
	return rdr.client.rdb.HSet(ctx, fmt.Sprintf("driver:%d:state", driver_id), "state", int(state)).Err()
}

func (rdr *RedisDriverRepository) GetDriverState(ctx context.Context, driver_id int) (models.DriverState, error) {
	res, err := rdr.client.rdb.HGet(ctx, fmt.Sprintf("driver:%d:state", driver_id), "state").Result()
	if err == redis.Nil {
		return models.IDLE, nil
	}
	if err != nil {
		return models.IDLE, err
	}
	val, err := strconv.Atoi(res)
	if err != nil {
		return models.IDLE, err
	}
	return models.DriverState(val), nil
}
