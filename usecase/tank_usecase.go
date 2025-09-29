package usecase

import "wot-statistics-server/domain"

// TankUseCase defines the interface for tank-related business logic.

type tankUseCase struct{}

func NewTankUseCase() domain.TankUseCase {
	return &tankUseCase{}
}

func (u *tankUseCase) GetTanks() (domain.Tank, error) {
	return domain.Tank{ID: "1", Name: "Lion", Type: "Medium", Tier: "X"}, nil
}
