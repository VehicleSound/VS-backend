package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/timickb/transport-sound/internal/config"
	"github.com/timickb/transport-sound/internal/controller"
	"github.com/timickb/transport-sound/internal/delivery"
	"github.com/timickb/transport-sound/internal/repository"
	"github.com/timickb/transport-sound/internal/usecase"
	"io/ioutil"
	"log"
)

func main() {
	cfg, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	connStr := "user=timickb dbname=soundp sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewPqRepository(db)

	userS := usecase.NewUserUseCase(repo)
	authS := usecase.NewAuthUseCase(repo)

	authC := controller.NewAuthController(authS, cfg.Secret)
	userC := controller.NewUserController(userS)

	s := delivery.NewHttpServer(authC, userC)

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
