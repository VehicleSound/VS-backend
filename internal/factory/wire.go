package factory

import (
	"database/sql"
	"fmt"
	"github.com/timickb/transport-sound/internal/config"
	"github.com/timickb/transport-sound/internal/controller"
	"github.com/timickb/transport-sound/internal/delivery/http/v1"
	"github.com/timickb/transport-sound/internal/infrastructure/repository/postgres"
	"github.com/timickb/transport-sound/internal/infrastructure/storage"
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
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.Name,
		cfg.Postgres.SslMode,
		cfg.Postgres.Port,
		cfg.Postgres.Password)

	// Init postgres
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Init minio
	store, err := storage.New(logger, storage.Params{
		Endpoint:        cfg.Minio.Endpoint,
		AccessKey:       cfg.Minio.AccessKey,
		SecretKey:       cfg.Minio.SecretKey,
		RequiredBuckets: []string{"images", "sounds"},
	})

	if err != nil {
		return nil, err
	}

	repo := postgres.NewPqRepository(db)

	userService := user.New(repo, logger)
	authService := auth.New(repo, logger)
	tagService := tag.New(repo, logger)
	soundService := sound.New(repo, logger)
	fileService := file.New(repo, store)
	searchService := search.New(repo, logger)

	authController := controller.NewAuth(authService, cfg.JwtSecret)
	userController := controller.NewUser(userService)
	tagController := controller.NewTag(tagService)
	soundController := controller.NewSound(soundService)
	fileController := controller.NewFile(fileService)
	searchController := controller.NewSearch(searchService)

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
