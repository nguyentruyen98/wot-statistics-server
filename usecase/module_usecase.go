package usecase

import (
	"encoding/json"
	"fmt"
	"os"
	"wot-statistics-server/domain"
)

type moduleUseCase struct {
	moduleRepo  domain.ModuleRepository
	profileRepo domain.ProfileRepository
}

// NewModuleUseCase creates a new instance of ModuleUseCase
func NewModuleUseCase(moduleRepo domain.ModuleRepository, profileRepo domain.ProfileRepository) domain.ModuleUseCase {
	return &moduleUseCase{
		moduleRepo:  moduleRepo,
		profileRepo: profileRepo,
	}
}

// GetModules retrieves all modules from database
func (u *moduleUseCase) GetModules() ([]domain.Module, error) {
	return u.moduleRepo.GetModules()
}

// GetModuleByID retrieves a single module by its module_id
func (u *moduleUseCase) GetModuleByID(moduleID int) (*domain.Module, error) {
	if moduleID <= 0 {
		return nil, fmt.Errorf("invalid module_id: must be greater than 0")
	}
	return u.moduleRepo.GetModuleByID(moduleID)
}

// GetModulesByFilter retrieves modules with filters
func (u *moduleUseCase) GetModulesByFilter(moduleType string, nation string, tier int) ([]domain.Module, error) {
	// Validate tier if provided
	if tier < 0 || tier > 10 {
		return nil, fmt.Errorf("invalid tier: must be between 1 and 10")
	}

	// Validate module type if provided
	validTypes := map[string]bool{
		"vehicleChassis": true,
		"vehicleTurret":  true,
		"vehicleGun":     true,
		"vehicleEngine":  true,
		"vehicleRadio":   true,
	}
	if moduleType != "" && !validTypes[moduleType] {
		return nil, fmt.Errorf("invalid module type: must be one of vehicleChassis, vehicleTurret, vehicleGun, vehicleEngine, vehicleRadio")
	}

	return u.moduleRepo.GetModulesByFilter(moduleType, nation, tier)
}

// ImportModulesFromJSON imports modules from JSON file
func (u *moduleUseCase) ImportModulesFromJSON(filePath string) (int, error) {
	// Read JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Parse JSON structure
	var apiResponse struct {
		Status string                          `json:"status"`
		Meta   map[string]interface{}          `json:"meta"`
		Data   map[string]domain.ModuleFromAPI `json:"data"`
	}

	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return 0, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Validate status
	if apiResponse.Status != "ok" {
		return 0, fmt.Errorf("API response status is not ok: %s", apiResponse.Status)
	}

	// Convert to domain modules
	modules := make([]*domain.Module, 0, len(apiResponse.Data))
	profiles := make([]*domain.Profile, 0, len(apiResponse.Data))

	for _, moduleAPI := range apiResponse.Data {
		module := moduleAPI.ToDomain()
		modules = append(modules, module)

		// Parse and create profile if default_profile exists
		if moduleAPI.DefaultProfile != nil {
			profile := domain.ModuleProfileFromDefaultProfile(moduleAPI.ModuleID, moduleAPI.DefaultProfile)
			if profile != nil {
				profiles = append(profiles, profile)
			}
		}
	}

	if len(modules) == 0 {
		return 0, fmt.Errorf("no modules found in JSON file")
	}

	// Insert modules to database
	count, err := u.moduleRepo.CreateModules(modules)
	if err != nil {
		return count, fmt.Errorf("failed to insert modules: %w", err)
	}

	// Insert profiles
	if len(profiles) > 0 {
		for _, profile := range profiles {
			if err := u.profileRepo.CreateProfile(profile); err != nil {
				// Log error but don't fail the import
				fmt.Printf("Warning: failed to create profile for module %v: %v\n", profile.ModuleID, err)
			}
		}
	}

	return count, nil
}

// CreateModules creates multiple modules in database (bulk insert)
func (u *moduleUseCase) CreateModules(modules []*domain.Module) (int, error) {
	if len(modules) == 0 {
		return 0, fmt.Errorf("no modules provided")
	}

	// Validate each module
	for i, module := range modules {
		if module.ModuleID <= 0 {
			return 0, fmt.Errorf("module %d has invalid module_id: must be greater than 0", i)
		}
		if module.Name == "" {
			return 0, fmt.Errorf("module %d has empty name", i)
		}
		if module.Tier < 1 || module.Tier > 10 {
			return 0, fmt.Errorf("module %d has invalid tier: must be between 1 and 10", i)
		}
	}

	return u.moduleRepo.CreateModules(modules)
}
