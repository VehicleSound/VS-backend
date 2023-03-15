package factory

import (
	"database/sql"
	"fmt"
	"github.com/timickb/transport-sound/internal/config"
	"github.com/timickb/transport-sound/internal/controller"
	"github.com/timickb/transport-sound/internal/delivery/http"
	"github.com/timickb/transport-sound/internal/interfaces"
	"github.com/timickb/transport-sound/internal/repository/postgres"
	"github.com/timickb/transport-sound/internal/usecase"
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

	repo := postgres.NewPqRepository(db)

	userService := usecase.NewUserUseCase(repo, logger)
	authService := usecase.NewAuthUseCase(repo, logger)
	tagService := usecase.NewTagUseCase(repo, logger)
	soundService := usecase.NewSoundUseCase(repo, logger)
	fileService := usecase.NewFileUseCase(repo, logger, cfg.MaxSoundSize)
	searchService := usecase.NewSearchUseCase(repo, logger)

	authController := controller.NewAuthController(authService, cfg.JwtSecret)
	userController := controller.NewUserController(userService)
	tagController := controller.NewTagController(tagService)
	soundController := controller.NewSoundController(soundService)
	fileController := controller.NewFileController(fileService)
	searchController := controller.NewSearchController(searchService)

	return http.NewHttpServer(
		cfg,
		authController,
		userController,
		tagController,
		soundController,
		fileController,
		searchController), nil
}
