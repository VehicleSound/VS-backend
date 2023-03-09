package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/timickb/transport-sound/internal/config"
	"github.com/timickb/transport-sound/internal/controller"
	"github.com/timickb/transport-sound/internal/delivery"
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
		"user=%s dbname=%s sslmode=%s password=%s",
		cfg.DbUser,
		cfg.DbName,
		cfg.DbSslMode,
		cfg.DbPassword)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	repo := postgres.NewPqRepository(db)

	userS := usecase.NewUserUseCase(repo)
	authS := usecase.NewAuthUseCase(repo)
	tagS := usecase.NewTagUseCase(repo)

	authC := controller.NewAuthController(authS, cfg.Secret)
	userC := controller.NewUserController(userS)
	tagC := controller.NewTagController(tagS)
	s := delivery.NewHttpServer(authC, userC, tagC)

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
