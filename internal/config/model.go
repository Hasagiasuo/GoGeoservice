package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresConfig PostgresConfig `json:"postgres"`
	RedisConfig    RedisConfig    `json:"redis"`
}

func UploadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("cannot load dotenv: %v", err)
		return nil
	}
	path_to_config := os.Getenv("config_path")
	if len(path_to_config) <= 0 {
		log.Fatal("path to config is empty")
		return nil
	}
	buffer, err := os.ReadFile(path_to_config)
	if err != nil {
		log.Fatal("cannot read config file")
		return nil
	}
	var cfg Config
	if err := json.Unmarshal(buffer, &cfg); err != nil {
		log.Fatal("cannot parse json from config file")
		return nil
	}
	return &cfg
}
