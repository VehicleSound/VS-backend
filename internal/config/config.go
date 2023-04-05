package config

type postgres struct {
	Host     string `json:"host,omitempty"`
	Name     string `json:"name,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	SslMode  string `json:"ssl_mode"`
	Port     int    `json:"port,omitempty"`
}

type kafka struct {
	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`
}

type AppConfig struct {
	AppPort        int    `json:"app_port,omitempty"`
	AppMetricsPort int    `json:"metrics_port,omitempty"`
	ServerMode     string `json:"server_mode,omitempty"`
	JwtSecret      string `json:"jwt_secret,omitempty"`
	MaxSoundSize   int    `json:"max_sound_size,omitempty"`
	MaxPictureSize int    `json:"max_picture_size,omitempty"`

	Postgres postgres `json:"postgres,omitempty"`
	Kafka    kafka    `json:"kafka,omitempty"`
}

func NewDefault() *AppConfig {
	return &AppConfig{
		AppPort:        8080,
		AppMetricsPort: 8081,
		JwtSecret:      "very_secured_secret",
		MaxSoundSize:   25242880,
		MaxPictureSize: 25242880,
		ServerMode:     "debug",
		Postgres: postgres{
			Host:     "localhost",
			Name:     "soundp",
			User:     "soundp",
			Password: "",
			SslMode:  "disable",
			Port:     5432,
		},
		Kafka: kafka{
			Host: "localhost",
			Port: 9092,
		},
	}
}
