package main

import (
	_ "github.com/lib/pq"
	"github.com/timickb/transport-sound/internal/config"
	"github.com/timickb/transport-sound/internal/factory"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := mainNoExit(); err != nil {
		log.Fatal(err)
	}
}

func mainNoExit() error {
	cfg := config.NewDefault()
	parseConfigFromEnvironment(cfg)

	srv, err := factory.InitializeHttpServer(cfg)
	if err != nil {
		return err
	}

	if err := srv.Run(); err != nil {
		return err
	}

	return nil
}

func parseConfigFromEnvironment(cfg *config.AppConfig) {
	if os.Getenv("DB_HOST") != "" {
		cfg.DbHost = os.Getenv("DB_HOST")
	}
	if os.Getenv("DB_USER") != "" {
		cfg.DbUser = os.Getenv("DB_USER")
	}
	if os.Getenv("DB_NAME") != "" {
		cfg.DbName = os.Getenv("DB_NAME")
	}
	if os.Getenv("DB_PASSWORD") != "" {
		cfg.DbPassword = os.Getenv("DB_PASSWORD")
	}
	if os.Getenv("JWT_SECRET") != "" {
		cfg.JwtSecret = os.Getenv("JWT_SECRET")
	}
	if os.Getenv("DB_PORT") != "" {
		cfg.DbPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	}
	if os.Getenv("APP_PORT") != "" {
		cfg.AppPort, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	}
	if os.Getenv("MAX_PICTURE_SIZE") != "" {
		cfg.MaxPictureSize, _ = strconv.Atoi(os.Getenv("MAX_PICTURE_SIZE"))
	}
	if os.Getenv("MAX_SOUND_SIZE") != "" {
		cfg.MaxSoundSize, _ = strconv.Atoi(os.Getenv("MAX_SOUND_SIZE"))
	}
}
