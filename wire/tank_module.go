//go:build wireinject
// +build wireinject

package wire

import (
	"wot-statistics-server/api/controller"
	"wot-statistics-server/config"
	"wot-statistics-server/internal/infrastructure/postgres"
	"wot-statistics-server/repository"
	"wot-statistics-server/usecase"

	"github.com/google/wire"
)

var TankSet = wire.NewSet(
	repository.NewTankRepository,
	repository.NewProfileRepository,
	usecase.NewTankUseCase,
	controller.NewTankController,
)

func InitializeTankController(db *postgres.DB, appConfig *config.AppConfig) *controller.TankController {
	wire.Build(TankSet)
	return nil
}
