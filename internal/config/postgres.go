package config

type PostgresConfig struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	User    string `json:"user"`
	DbName  string `json:"db_name"`
	SslMode string `json:"ssl_mode"`
}
