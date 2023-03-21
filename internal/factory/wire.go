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

	userService := user.NewUserUseCase(repo, logger)
	authService := auth.NewAuthUseCase(repo, logger)
	tagService := tag.NewTagUseCase(repo, logger)
	soundService := sound.NewSoundUseCase(repo, logger)
	fileService := file.NewFileUseCase(repo, logger, cfg.MaxSoundSize)
	searchService := search.NewSearchUseCase(repo, logger)

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
