//go:build wireinject
// +build wireinject

package wire

import (
	"wot-statistics-server/api/controller"
	"wot-statistics-server/config"
	"wot-statistics-server/repository"
	"wot-statistics-server/usecase"

	"github.com/google/wire"
)

var TankSet = wire.NewSet(repository.NewTankRepository, usecase.NewTankUseCase, controller.NewTankController)

func InitializeTankController(appConfig *config.AppConfig) *controller.TankController {
	wire.Build(TankSet)
	return nil
}
