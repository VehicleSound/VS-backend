//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/google/wire"
	"github.com/timickb/transport-sound/internal/controller"
	"github.com/timickb/transport-sound/internal/delivery"
	"github.com/timickb/transport-sound/internal/repository/postgres"
	"github.com/timickb/transport-sound/internal/usecase"
)

func InitializeHttpServer() *delivery.HttpServer {
	wire.Build(
		postgres.NewPqRepository,
		usecase.NewUserUseCase,
		usecase.NewTagUseCase,
		usecase.NewAuthUseCase,
		controller.NewAuthController,
		controller.NewUserController,
		controller.NewTagController,
		delivery.NewHttpServer)
	return &delivery.HttpServer{}
}
