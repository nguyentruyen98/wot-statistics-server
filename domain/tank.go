package domain

import "wot-statistics-server/external"

type Tank struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Tier string `json:"tier"`
}
type TankUseCase interface {
	// GetTanks retrieves a list of tanks.
	GetTanks() ([]Tank, error)
	FetchTanksFromAPI() (*external.APIResponse, error)
}

type TankRepository interface {
	GetTanks() ([]Tank, error)
	GetTanksFromAPI() (*external.APIResponse, error)
}
