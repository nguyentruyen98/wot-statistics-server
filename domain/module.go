package domain

import "time"

// Module represents a World of Tanks vehicle module (component)
type Module struct {
	ID          int       `json:"id" db:"id"`
	ModuleID    int       `json:"module_id" db:"module_id"`
	Name        string    `json:"name" db:"name"`
	Type        string    `json:"type" db:"type"`
	Nation      string    `json:"nation" db:"nation"`
	Tier        int       `json:"tier" db:"tier"`
	Weight      int       `json:"weight" db:"weight"`
	PriceCredit *int      `json:"price_credit" db:"price_credit"`
	Image       string    `json:"image" db:"image"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// ModuleFromAPI represents the raw module data from Wargaming API
type ModuleFromAPI struct {
	ModuleID       int                    `json:"module_id"`
	Name           string                 `json:"name"`
	Type           string                 `json:"type"`
	Nation         string                 `json:"nation"`
	Tier           int                    `json:"tier"`
	Weight         int                    `json:"weight"`
	PriceCredit    *int                   `json:"price_credit"`
	Image          string                 `json:"image"`
	DefaultProfile map[string]interface{} `json:"default_profile"`
}

// ToDomain converts ModuleFromAPI to domain Module
func (m *ModuleFromAPI) ToDomain() *Module {
	return &Module{
		ModuleID:    m.ModuleID,
		Name:        m.Name,
		Type:        m.Type,
		Nation:      m.Nation,
		Tier:        m.Tier,
		Weight:      m.Weight,
		PriceCredit: m.PriceCredit,
		Image:       m.Image,
	}
}

// ModuleUseCase defines business logic operations for modules
type ModuleUseCase interface {
	// GetModules retrieves all modules from database
	GetModules() ([]Module, error)

	// GetModuleByID retrieves a single module by its module_id
	GetModuleByID(moduleID int) (*Module, error)

	// GetModulesByFilter retrieves modules with filters
	GetModulesByFilter(moduleType string, nation string, tier int) ([]Module, error)

	// ImportModulesFromJSON imports modules from JSON file
	ImportModulesFromJSON(filePath string) (int, error)

	// CreateModules creates multiple modules in database (bulk insert)
	CreateModules(modules []*Module) (int, error)
}

// ModuleRepository defines data access operations for modules
type ModuleRepository interface {
	// GetModules retrieves all modules from database
	GetModules() ([]Module, error)

	// GetModuleByID retrieves a single module by its module_id
	GetModuleByID(moduleID int) (*Module, error)

	// GetModulesByFilter retrieves modules with filters
	GetModulesByFilter(moduleType string, nation string, tier int) ([]Module, error)

	// CreateModules inserts multiple modules into database (bulk insert)
	CreateModules(modules []*Module) (int, error)

	// ModuleExists checks if a module exists by module_id
	ModuleExists(moduleID int) (bool, error)
}
