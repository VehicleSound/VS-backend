package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	log2 "github.com/qiniu/x/log"
	"github.com/timickb/transport-sound/internal/config"
	"github.com/timickb/transport-sound/internal/controller"
	"github.com/timickb/transport-sound/internal/delivery/http"
	"github.com/timickb/transport-sound/internal/repository/postgres"
	"github.com/timickb/transport-sound/internal/usecase"
	"io/ioutil"
	"log"
)

func main() {
	cfg, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s port=%d password=%s",
		cfg.DbHost,
		cfg.DbUser,
		cfg.DbName,
		cfg.DbSslMode,
		cfg.DbPort,
		cfg.DbPassword)

	log2.Info(fmt.Sprintf("Connection string: postgresql://%s:%s@%s/%s?sslmode=%s",
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbName,
		cfg.DbSslMode))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	repo := postgres.NewPqRepository(db)

	userS := usecase.NewUserUseCase(repo)
	authS := usecase.NewAuthUseCase(repo)
	tagS := usecase.NewTagUseCase(repo)
	soundS := usecase.NewSoundUseCase(repo)
	fileS := usecase.NewFileUseCase(repo, cfg.MaxSoundSize)
	searchS := usecase.NewSearchUseCase(repo)

	authC := controller.NewAuthController(authS, cfg.Secret)
	userC := controller.NewUserController(userS)
	tagC := controller.NewTagController(tagS)
	soundC := controller.NewSoundController(soundS)
	fileC := controller.NewFileController(fileS)
	searchC := controller.NewSearchController(searchS)

	s := http.NewHttpServer(cfg, authC, userC, tagC, soundC, fileC, searchC)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}

}

func ReadConfig() (*config.Config, error) {
	raw, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var cfg *config.Config

	if json.Unmarshal(raw, &cfg) != nil {
		return nil, err
	}

	return cfg, nil
}
