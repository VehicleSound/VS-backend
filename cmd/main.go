package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/timickb/transport-sound/internal/config"
	"github.com/timickb/transport-sound/internal/factory"
	"github.com/timickb/transport-sound/internal/interfaces"
	"github.com/timickb/transport-sound/internal/metrics"
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

	mts := metrics.New(logger, cfg.AppMetricsPort)

	go func() {
		logger.Info(fmt.Sprintf("Starting metrics listener on port %d", cfg.AppMetricsPort))
		if err := mts.Listen(); err != nil {
			logger.Fatal(err)
		}
	}()

	logger.Info("Initializing http server")

	srv, err := factory.InitializeHttpServer(cfg, logger, mts)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Starting http server on port %d", cfg.AppPort))
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
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
	if os.Getenv("APP_METRICS_PORT") != "" {
		cfg.AppMetricsPort, _ = strconv.Atoi(os.Getenv("APP_METRICS_PORT"))
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
