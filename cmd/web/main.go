package main

import (
	"fmt"
	"geoserv/internal/config"
	"geoserv/internal/infrastructure/postgres"
	"geoserv/internal/infrastructure/redis"
	"geoserv/internal/usecase"
)

func main() {
	cfg := config.UploadConfig()
	client := postgres.NewPostgresClient(cfg.PostgresConfig)
	rc := redis.NewRedisClient(cfg.RedisConfig)
	dpr := redis.NewRedisDriverPositionRepository(rc)
	dr := postgres.NewPostgresDriverRepository(client)
	du := usecase.NewDriverUsecase(dr, dpr)
	drivers, err := du.GetAll()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(drivers)
}
