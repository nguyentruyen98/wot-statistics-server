package usecase

import (
	"wot-statistics-server/domain"
	"wot-statistics-server/external"
)

type tankUseCase struct {
	tankRepository domain.TankRepository
}

func NewTankUseCase(tankRepository domain.TankRepository) domain.TankUseCase {
	return &tankUseCase{
		tankRepository: tankRepository,
	}
}

func (u *tankUseCase) GetTanks() ([]domain.Tank, error) {
	return u.tankRepository.GetTanks()
}

func (u *tankUseCase) FetchTanksFromAPI() (*external.APIResponse, error) {
	return u.tankRepository.GetTanksFromAPI()
}
