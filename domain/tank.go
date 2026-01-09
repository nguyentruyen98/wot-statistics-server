package domain

import (
	"database/sql/driver"
	"encoding/json"
	"time"
	"wot-statistics-server/external"
)

// JSONB type for PostgreSQL JSONB columns
type JSONB map[string]interface{}

// Value implements the driver.Valuer interface for JSONB
func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// Scan implements the sql.Scanner interface for JSONB
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, j)
}

// IntArray type for PostgreSQL integer arrays stored as JSONB
type IntArray []int

// Value implements the driver.Valuer interface
func (ia IntArray) Value() (driver.Value, error) {
	if ia == nil {
		return nil, nil
	}
	return json.Marshal(ia)
}

// Scan implements the sql.Scanner interface
func (ia *IntArray) Scan(value interface{}) error {
	if value == nil {
		*ia = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, ia)
}

// Tank represents a World of Tanks vehicle
type Tank struct {
	ID        int    `json:"id" db:"id"`
	TankID    int    `json:"tank_id" db:"tank_id"`
	Name      string `json:"name" db:"name"`
	ShortName string `json:"short_name" db:"short_name"`
	Nation    string `json:"nation" db:"nation"`
	Tier      int    `json:"tier" db:"tier"`
	Type      string `json:"type" db:"type"`
	Tag       string `json:"tag" db:"tag"`

	// Premium and gift status
	IsPremium    bool `json:"is_premium" db:"is_premium"`
	IsPremiumIGR bool `json:"is_premium_igr" db:"is_premium_igr"`
	IsGift       bool `json:"is_gift" db:"is_gift"`
	IsWheeled    bool `json:"is_wheeled" db:"is_wheeled"`

	// Pricing (can be null for premium/gift tanks)
	PriceCredit *int64 `json:"price_credit" db:"price_credit"`
	PriceGold   *int   `json:"price_gold" db:"price_gold"`

	// Description and images
	Description string `json:"description" db:"description"`
	BigIconURL  string `json:"big_icon_url" db:"big_icon_url"`

	// Modules (stored as JSONB arrays)
	Radios      IntArray `json:"radios" db:"radios"`
	Suspensions IntArray `json:"suspensions" db:"suspensions"`
	Provisions  IntArray `json:"provisions" db:"provisions"`
	Engines     IntArray `json:"engines" db:"engines"`
	Guns        IntArray `json:"guns" db:"guns"`
	Turrets     IntArray `json:"turrets" db:"turrets"`
	Crew        JSONB    `json:"crew" db:"crew"`

	// Technical specifications
	ModulesTree JSONB `json:"modules_tree" db:"modules_tree"`
	NextTanks   JSONB `json:"next_tanks" db:"next_tanks"`
	PricesXP    JSONB `json:"prices_xp" db:"prices_xp"`

	// Metadata
	Multination *string   `json:"multination" db:"multination"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// TankFromAPI represents the raw tank data from Wargaming API
type TankFromAPI struct {
	TankID         int                    `json:"tank_id"`
	Name           string                 `json:"name"`
	ShortName      string                 `json:"short_name"`
	Nation         string                 `json:"nation"`
	Tier           int                    `json:"tier"`
	Type           string                 `json:"type"`
	Tag            string                 `json:"tag"`
	IsPremium      bool                   `json:"is_premium"`
	IsPremiumIGR   bool                   `json:"is_premium_igr"`
	IsGift         bool                   `json:"is_gift"`
	IsWheeled      bool                   `json:"is_wheeled"`
	PriceCredit    *int64                 `json:"price_credit"`
	PriceGold      *int                   `json:"price_gold"`
	Description    string                 `json:"description"`
	Images         map[string]string      `json:"images"`
	Radios         []int                  `json:"radios"`
	Suspensions    []int                  `json:"suspensions"`
	Provisions     []int                  `json:"provisions"`
	Engines        []int                  `json:"engines"`
	Guns           []int                  `json:"guns"`
	Turrets        []int                  `json:"turrets"`
	Crew           []interface{}          `json:"crew"`
	DefaultProfile map[string]interface{} `json:"default_profile"`
	ModulesTree    map[string]interface{} `json:"modules_tree"`
	NextTanks      map[string]interface{} `json:"next_tanks"`
	PricesXP       map[string]interface{} `json:"prices_xp"`
	Multination    *string                `json:"multination"`
}

// ToDomain converts TankFromAPI to domain Tank
func (t *TankFromAPI) ToDomain() *Tank {
	tank := &Tank{
		TankID:       t.TankID,
		Name:         t.Name,
		ShortName:    t.ShortName,
		Nation:       t.Nation,
		Tier:         t.Tier,
		Type:         t.Type,
		Tag:          t.Tag,
		IsPremium:    t.IsPremium,
		IsPremiumIGR: t.IsPremiumIGR,
		IsGift:       t.IsGift,
		IsWheeled:    t.IsWheeled,
		PriceCredit:  t.PriceCredit,
		PriceGold:    t.PriceGold,
		Description:  t.Description,
		Radios:       t.Radios,
		Suspensions:  t.Suspensions,
		Provisions:   t.Provisions,
		Engines:      t.Engines,
		Guns:         t.Guns,
		Turrets:      t.Turrets,
		Multination:  t.Multination,
	}

	// Extract big_icon URL from images
	if bigIcon, ok := t.Images["big_icon"]; ok {
		tank.BigIconURL = bigIcon
	}

	// Convert crew to JSONB
	if t.Crew != nil {
		tank.Crew = make(JSONB)
		crewBytes, _ := json.Marshal(t.Crew)
		json.Unmarshal(crewBytes, &tank.Crew)
	}

	// Convert modules_tree to JSONB
	if t.ModulesTree != nil {
		tank.ModulesTree = JSONB(t.ModulesTree)
	}

	// Convert next_tanks to JSONB
	if t.NextTanks != nil {
		tank.NextTanks = JSONB(t.NextTanks)
	}

	// Convert prices_xp to JSONB
	if t.PricesXP != nil {
		tank.PricesXP = JSONB(t.PricesXP)
	}

	return tank
}

// TankUseCase defines business logic operations for tanks
type TankUseCase interface {
	// GetTanks retrieves a list of tanks from database
	GetTanks() ([]Tank, error)

	// GetTankByID retrieves a single tank by its tank_id
	GetTankByID(tankID int) (*Tank, error)

	// GetTanksByFilter retrieves tanks with filters
	GetTanksByFilter(nation string, tier int, tankType string, isPremium *bool) ([]Tank, error)

	// GetTanksFromAPI fetches tanks from Wargaming API
	GetTanksFromAPI() (*external.APIResponse, error)

	// ImportTanksFromJSON imports tanks from JSON file
	ImportTanksFromJSON(filePath string) (int, error)

	// CreateTank creates a new tank in database
	CreateTank(tank *Tank) error

	// CreateTanks creates multiple tanks in database (bulk insert)
	CreateTanks(tanks []*Tank) (int, error)

	// UpdateTank updates an existing tank
	UpdateTank(tank *Tank) error

	// DeleteTank deletes a tank by ID
	DeleteTank(tankID int) error
}

// TankRepository defines data access operations for tanks
type TankRepository interface {
	// GetTanks retrieves all tanks from database
	GetTanks() ([]Tank, error)

	// GetTankByID retrieves a single tank by its tank_id
	GetTankByID(tankID int) (*Tank, error)

	// GetTanksByFilter retrieves tanks with filters
	GetTanksByFilter(nation string, tier int, tankType string, isPremium *bool) ([]Tank, error)

	// GetTanksFromAPI fetches tanks from Wargaming API
	GetTanksFromAPI() (*external.APIResponse, error)

	// CreateTank inserts a new tank into database
	CreateTank(tank *Tank) error

	// CreateTanks inserts multiple tanks into database (bulk insert)
	CreateTanks(tanks []*Tank) (int, error)

	// UpdateTank updates an existing tank in database
	UpdateTank(tank *Tank) error

	// DeleteTank deletes a tank by tank_id
	DeleteTank(tankID int) error

	// TankExists checks if a tank exists by tank_id
	TankExists(tankID int) (bool, error)
}
