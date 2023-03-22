package factory

import (
	"database/sql"
	"fmt"
	"github.com/timickb/transport-sound/internal/config"
	"github.com/timickb/transport-sound/internal/delivery/http"
	"github.com/timickb/transport-sound/internal/infrastructure/controller"
	"github.com/timickb/transport-sound/internal/infrastructure/repository/postgres"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase/auth"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase/file"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase/search"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase/sound"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase/tag"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase/user"
	"github.com/timickb/transport-sound/internal/interfaces"
)

func InitializeHttpServer(cfg *config.AppConfig, logger interfaces.Logger) (*http.Server, error) {
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

	authController := controller.NewAuth(authService, cfg.JwtSecret)
	userController := controller.NewUser(userService)
	tagController := controller.NewTag(tagService)
	soundController := controller.NewSound(soundService)
	fileController := controller.NewFile(fileService)
	searchController := controller.NewSearch(searchService)

	return http.NewHttpServer(
		cfg,
		authController,
		userController,
		tagController,
		soundController,
		fileController,
		searchController), nil
}
