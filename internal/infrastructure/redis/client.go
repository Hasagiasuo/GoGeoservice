package redis

import (
	"geoserv/internal/config"

	red "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	rdb *red.Client
}

func NewRedisClient(cfg config.RedisConfig) *RedisClient {
	rdb := red.NewClient(&red.Options{
		Addr:     cfg.Host,
		Password: cfg.Password,
		DB:       cfg.Db,
	})
	return &RedisClient{rdb}
}
