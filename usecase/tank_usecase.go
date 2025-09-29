package usecase

import "wot-statistics-server/domain"

// TankUseCase defines the interface for tank-related business logic.

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
