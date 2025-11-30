package config

type RedisConfig struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}
