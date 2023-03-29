package factory

import (
	"database/sql"
	"fmt"
	"github.com/timickb/transport-sound/internal/config"
	controller2 "github.com/timickb/transport-sound/internal/controller"
	"github.com/timickb/transport-sound/internal/delivery/http/v1"
	"github.com/timickb/transport-sound/internal/infrastructure/repository/postgres"
	"github.com/timickb/transport-sound/internal/interfaces"
	"github.com/timickb/transport-sound/internal/usecase/auth"
	"github.com/timickb/transport-sound/internal/usecase/file"
	"github.com/timickb/transport-sound/internal/usecase/search"
	"github.com/timickb/transport-sound/internal/usecase/sound"
	"github.com/timickb/transport-sound/internal/usecase/tag"
	"github.com/timickb/transport-sound/internal/usecase/user"
)

func InitializeHttpServer(cfg *config.AppConfig, logger interfaces.Logger, metrics interfaces.Metrics) (*v1.Server, error) {
	connStr := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s port=%d password=%s",
		cfg.DbHost,
		cfg.DbUser,
		cfg.DbName,
		cfg.DbSslMode,
		cfg.DbPort,
		cfg.DbPassword)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	repo := postgres.NewPqRepository(db)

	userService := user.New(repo, logger)
	authService := auth.New(repo, logger)
	tagService := tag.New(repo, logger)
	soundService := sound.New(repo, logger)
	fileService := file.New(repo, logger, cfg.MaxSoundSize)
	searchService := search.New(repo, logger)

	authController := controller2.NewAuth(authService, cfg.JwtSecret)
	userController := controller2.NewUser(userService)
	tagController := controller2.NewTag(tagService)
	soundController := controller2.NewSound(soundService)
	fileController := controller2.NewFile(fileService)
	searchController := controller2.NewSearch(searchService)

	return v1.NewHttpServer(
		cfg,
		metrics,
		authController,
		userController,
		tagController,
		soundController,
		fileController,
		searchController), nil
}
