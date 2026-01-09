package usecase

import (
	"encoding/json"
	"fmt"
	"os"

	"wot-statistics-server/domain"
	"wot-statistics-server/external"
)

type tankUseCase struct {
	tankRepository    domain.TankRepository
	profileRepository domain.ProfileRepository
}

// NewTankUseCase creates a new tank use case instance
func NewTankUseCase(tankRepository domain.TankRepository, profileRepository domain.ProfileRepository) domain.TankUseCase {
	return &tankUseCase{
		tankRepository:    tankRepository,
		profileRepository: profileRepository,
	}
}

// GetTanks retrieves all tanks from database
func (u *tankUseCase) GetTanks() ([]domain.Tank, error) {
	return u.tankRepository.GetTanks()
}

// GetTankByID retrieves a single tank by its tank_id
func (u *tankUseCase) GetTankByID(tankID int) (*domain.Tank, error) {
	if tankID <= 0 {
		return nil, fmt.Errorf("invalid tank ID: %d", tankID)
	}
	return u.tankRepository.GetTankByID(tankID)
}

// GetTanksByFilter retrieves tanks with filters
func (u *tankUseCase) GetTanksByFilter(nation string, tier int, tankType string, isPremium *bool) ([]domain.Tank, error) {
	// Validate tier
	if tier < 0 || tier > 10 {
		return nil, fmt.Errorf("invalid tier: %d (must be between 1 and 10)", tier)
	}

	// Validate nation
	validNations := map[string]bool{
		"ussr": true, "germany": true, "usa": true, "france": true,
		"uk": true, "china": true, "japan": true, "czech": true,
		"sweden": true, "poland": true, "italy": true,
	}
	if nation != "" && !validNations[nation] {
		return nil, fmt.Errorf("invalid nation: %s", nation)
	}

	// Validate type
	validTypes := map[string]bool{
		"heavyTank": true, "mediumTank": true, "lightTank": true,
		"AT-SPG": true, "SPG": true,
	}
	if tankType != "" && !validTypes[tankType] {
		return nil, fmt.Errorf("invalid tank type: %s", tankType)
	}

	return u.tankRepository.GetTanksByFilter(nation, tier, tankType, isPremium)
}

// GetTanksFromAPI fetches tanks from Wargaming API
func (u *tankUseCase) GetTanksFromAPI() (*external.APIResponse, error) {
	return u.tankRepository.GetTanksFromAPI()
}

// ImportTanksFromJSON reads a JSON file and imports tanks into database
func (u *tankUseCase) ImportTanksFromJSON(filePath string) (int, error) {
	// Read file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse JSON structure
	var apiResponse struct {
		Status string                        `json:"status"`
		Meta   map[string]interface{}        `json:"meta"`
		Data   map[string]domain.TankFromAPI `json:"data"`
	}

	err = json.Unmarshal(fileData, &apiResponse)
	if err != nil {
		return 0, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Check status
	if apiResponse.Status != "ok" {
		return 0, fmt.Errorf("API response status is not ok: %s", apiResponse.Status)
	}

	// Convert to domain tanks
	tanks := make([]*domain.Tank, 0, len(apiResponse.Data))
	profiles := make([]*domain.Profile, 0, len(apiResponse.Data))

	for _, tankFromAPI := range apiResponse.Data {
		tank := tankFromAPI.ToDomain()
		tanks = append(tanks, tank)

		// Convert default_profile to Profile
		if tankFromAPI.DefaultProfile != nil {
			profile := domain.ProfileFromDefaultProfile(tankFromAPI.TankID, tankFromAPI.DefaultProfile)
			profiles = append(profiles, profile)
		}
	}

	if len(tanks) == 0 {
		return 0, fmt.Errorf("no tanks found in JSON file")
	}

	// Bulk insert tanks into database
	count, err := u.tankRepository.CreateTanks(tanks)
	if err != nil {
		return count, fmt.Errorf("failed to import tanks: %w", err)
	}

	// Bulk insert profiles into database
	if len(profiles) > 0 {
		successCount := 0
		failedCount := 0

		for _, profile := range profiles {
			err := u.profileRepository.CreateProfile(profile)
			if err != nil {
				fmt.Printf("Warning: failed to create profile for tank_id %d: %v\n", profile.TankID, err)
				failedCount++
			} else {
				successCount++
			}
		}

		fmt.Printf("Profiles: %d created successfully, %d failed\n", successCount, failedCount)
	}

	return count, nil
}

// CreateTank creates a new tank in database
func (u *tankUseCase) CreateTank(tank *domain.Tank) error {
	// Validate tank data
	if tank.TankID <= 0 {
		return fmt.Errorf("invalid tank_id: %d", tank.TankID)
	}
	if tank.Name == "" {
		return fmt.Errorf("tank name is required")
	}
	if tank.ShortName == "" {
		return fmt.Errorf("tank short_name is required")
	}
	if tank.Nation == "" {
		return fmt.Errorf("tank nation is required")
	}
	if tank.Tier < 1 || tank.Tier > 10 {
		return fmt.Errorf("invalid tier: %d (must be between 1 and 10)", tank.Tier)
	}
	if tank.Type == "" {
		return fmt.Errorf("tank type is required")
	}

	return u.tankRepository.CreateTank(tank)
}

// CreateTanks creates multiple tanks in database (bulk insert)
func (u *tankUseCase) CreateTanks(tanks []*domain.Tank) (int, error) {
	if len(tanks) == 0 {
		return 0, fmt.Errorf("no tanks to create")
	}

	// Validate all tanks before insertion
	for i, tank := range tanks {
		if tank.TankID <= 0 {
			return 0, fmt.Errorf("tank at index %d has invalid tank_id: %d", i, tank.TankID)
		}
		if tank.Name == "" {
			return 0, fmt.Errorf("tank at index %d has empty name", i)
		}
		if tank.Tier < 1 || tank.Tier > 10 {
			return 0, fmt.Errorf("tank at index %d has invalid tier: %d", i, tank.Tier)
		}
	}

	return u.tankRepository.CreateTanks(tanks)
}

// UpdateTank updates an existing tank
func (u *tankUseCase) UpdateTank(tank *domain.Tank) error {
	// Validate tank data
	if tank.TankID <= 0 {
		return fmt.Errorf("invalid tank_id: %d", tank.TankID)
	}

	// Check if tank exists
	exists, err := u.tankRepository.TankExists(tank.TankID)
	if err != nil {
		return fmt.Errorf("failed to check tank existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("tank with ID %d does not exist", tank.TankID)
	}

	return u.tankRepository.UpdateTank(tank)
}

// DeleteTank deletes a tank by ID
func (u *tankUseCase) DeleteTank(tankID int) error {
	if tankID <= 0 {
		return fmt.Errorf("invalid tank_id: %d", tankID)
	}

	// Check if tank exists before deleting
	exists, err := u.tankRepository.TankExists(tankID)
	if err != nil {
		return fmt.Errorf("failed to check tank existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("tank with ID %d does not exist", tankID)
	}

	return u.tankRepository.DeleteTank(tankID)
}
