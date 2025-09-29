package repository

import "wot-statistics-server/domain"

type tankRepository struct {
	tanks []domain.Tank
}

func NewTankRepository() domain.TankRepository {
	return &tankRepository{
		tanks: []domain.Tank{
			{ID: "1", Name: "Lion", Type: "Medium", Tier: "X"},
			{ID: "2", Name: "Tiger", Type: "Heavy", Tier: "VIII"},
		},
	}
}

func (r *tankRepository) GetTanks() ([]domain.Tank, error) {
	return r.tanks, nil
}
