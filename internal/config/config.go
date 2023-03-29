package config

type AppConfig struct {
	AppPort        int    `json:"app_port,omitempty"`
	AppMetricsPort int    `json:"metrics_port,omitempty"`
	JwtSecret      string `json:"jwt_secret,omitempty"`
	DbHost         string `json:"db_host,omitempty"`
	DbName         string `json:"db_name,omitempty"`
	DbUser         string `json:"db_user,omitempty"`
	DbPassword     string `json:"db_password,omitempty"`
	DbPort         int    `json:"db_port,omitempty"`
	DbSslMode      string `json:"db_ssl_mode,omitempty"`
	MaxSoundSize   int    `json:"max_sound_size,omitempty"`
	MaxPictureSize int    `json:"max_picture_size,omitempty"`
	ServerMode     string `json:"server_mode,omitempty"`
}

func NewDefault() *AppConfig {
	return &AppConfig{
		AppPort:        8080,
		AppMetricsPort: 8081,
		JwtSecret:      "very_secured_secret",
		DbHost:         "localhost",
		DbName:         "soundp",
		DbUser:         "soundp",
		DbPassword:     "",
		DbPort:         5432,
		DbSslMode:      "disable",
		MaxSoundSize:   25242880,
		MaxPictureSize: 25242880,
		ServerMode:     "debug",
	}
}
