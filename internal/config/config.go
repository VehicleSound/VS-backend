package config

type Config struct {
	Port       int    `json:"port,omitempty"`
	Secret     string `json:"secret,omitempty"`
	DbName     string `json:"db_name,omitempty"`
	DbUser     string `json:"db_user,omitempty"`
	DbPassword string `json:"db_password,omitempty"`
	DbSslMode  string `json:"db_ssl_mode,omitempty"`
}
