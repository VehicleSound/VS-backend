package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/timickb/transport-sound/internal/config"
	"github.com/timickb/transport-sound/internal/factory"
	"github.com/timickb/transport-sound/internal/interfaces"
	"os"
	"strconv"
)

func main() {
	logger := logrus.New()
	logger.Info("Service is starting")

	if err := mainNoExit(logger); err != nil {
		logger.Fatal(err)
	}
}

func mainNoExit(logger interfaces.Logger) error {
	cfg := config.NewDefault()
	parseConfigFromEnvironment(cfg)

	if cfg.ServerMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else if cfg.ServerMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		return errors.New("invalid server mode: it can be only \"debug\" or \"release\"")
	}

	logger.Info("Config parsed: ", cfg)
	logger.Info("Initializing http server")

	srv, err := factory.InitializeHttpServer(cfg)
	if err != nil {
		return err
	}

	logger.Info("Starting http server")
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
	if os.Getenv("SERVER_MODE") != "" {
		cfg.ServerMode = os.Getenv("SERVER_MODE")
	}
}
