package repository

import (
	"wot-statistics-server/domain"
	"wot-statistics-server/external"
)

type tankRepository struct {
	tanks        []domain.Tank
	wargamingAPI *external.WargamingAPI
}

func NewTankRepository(appID string) domain.TankRepository {
	return &tankRepository{
		tanks: []domain.Tank{
			{ID: "1", Name: "Lion", Type: "Medium", Tier: "X"},
			{ID: "2", Name: "Tiger", Type: "Heavy", Tier: "VIII"},
		},
		wargamingAPI: external.NewWargamingAPI(appID),
	}
}

func (r *tankRepository) GetTanks() ([]domain.Tank, error) {
	// For now, return static data
	// You can extend this to fetch from Wargaming API using r.wargamingAPI

	return r.tanks, nil
}

// GetTanksFromAPI fetches tanks from Wargaming API
func (r *tankRepository) GetTanksFromAPI() (*external.APIResponse, error) {
	params := map[string]string{
		"limit":  "100",
		"fields": "tank_id,name,type,tier",
	}
	return r.wargamingAPI.GetTanksList(params)
}
