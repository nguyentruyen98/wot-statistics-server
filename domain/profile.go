package domain

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// AmmoItem represents a shell type with its characteristics
type AmmoItem struct {
	Damage      []int  `json:"damage"`      // [min, avg, max]
	Penetration []int  `json:"penetration"` // [min, avg, max]
	Type        string `json:"type"`
	Stun        *Stun  `json:"stun,omitempty"`
}

// Stun represents stun characteristics for certain shell types
type Stun struct {
	Duration map[string]float64 `json:"duration"` // {"min": x, "max": y}
}

// RapidMode represents vehicle characteristics in Rapid mode (for wheeled vehicles)
type RapidMode struct {
	SpeedBackward               float64 `json:"speed_backward"`
	SpeedForward                float64 `json:"speed_forward"`
	SuspensionSteeringLockAngle int     `json:"suspension_steering_lock_angle"`
	SwitchOffTime               float64 `json:"switch_off_time"`
	SwitchOnTime                float64 `json:"switch_on_time"`
}

// SiegeMode represents vehicle characteristics in Siege mode
type SiegeMode struct {
	AimTime                 float64 `json:"aim_time"`
	Dispersion              float64 `json:"dispersion"`
	MoveDownArc             int     `json:"move_down_arc"`
	MoveUpArc               int     `json:"move_up_arc"`
	ReloadTime              float64 `json:"reload_time"`
	SpeedBackward           float64 `json:"speed_backward"`
	SuspensionTraverseSpeed int     `json:"suspension_traverse_speed"`
	SwitchOffTime           float64 `json:"switch_off_time"`
	SwitchOnTime            float64 `json:"switch_on_time"`
}

// Profile represents the default profile/configuration of a tank or module
type Profile struct {
	ID       int  `json:"id" db:"id"`
	TankID   *int `json:"tank_id,omitempty" db:"tank_id"`
	ModuleID *int `json:"module_id,omitempty" db:"module_id"`

	// Basic characteristics
	HP            *int     `json:"hp" db:"hp"`
	HullHP        *int     `json:"hull_hp" db:"hull_hp"`
	HullWeight    *int     `json:"hull_weight" db:"hull_weight"`
	MaxAmmo       *int     `json:"max_ammo" db:"max_ammo"`
	MaxWeight     *int     `json:"max_weight" db:"max_weight"`
	SpeedBackward *float64 `json:"speed_backward" db:"speed_backward"`
	SpeedForward  *float64 `json:"speed_forward" db:"speed_forward"`
	Weight        *int     `json:"weight" db:"weight"`

	// Ammo characteristics (stored as JSONB array)
	Ammo JSONB `json:"ammo" db:"ammo"`

	// Armor characteristics - Hull
	ArmorHullFront *int `json:"armor_hull_front" db:"armor_hull_front"`
	ArmorHullRear  *int `json:"armor_hull_rear" db:"armor_hull_rear"`
	ArmorHullSides *int `json:"armor_hull_sides" db:"armor_hull_sides"`

	// Armor characteristics - Turret
	ArmorTurretFront *int `json:"armor_turret_front" db:"armor_turret_front"`
	ArmorTurretRear  *int `json:"armor_turret_rear" db:"armor_turret_rear"`
	ArmorTurretSides *int `json:"armor_turret_sides" db:"armor_turret_sides"`

	// Engine characteristics
	EngineFireChance *float64 `json:"engine_fire_chance" db:"engine_fire_chance"`
	EngineName       *string  `json:"engine_name" db:"engine_name"`
	EnginePower      *int     `json:"engine_power" db:"engine_power"`
	EngineTag        *string  `json:"engine_tag" db:"engine_tag"`
	EngineTier       *int     `json:"engine_tier" db:"engine_tier"`
	EngineWeight     *int     `json:"engine_weight" db:"engine_weight"`

	// Gun characteristics
	GunAimTime       *float64 `json:"gun_aim_time" db:"gun_aim_time"`
	GunCaliber       *int     `json:"gun_caliber" db:"gun_caliber"`
	GunDispersion    *float64 `json:"gun_dispersion" db:"gun_dispersion"`
	GunFireRate      *float64 `json:"gun_fire_rate" db:"gun_fire_rate"`
	GunMoveDownArc   *int     `json:"gun_move_down_arc" db:"gun_move_down_arc"`
	GunMoveUpArc     *int     `json:"gun_move_up_arc" db:"gun_move_up_arc"`
	GunName          *string  `json:"gun_name" db:"gun_name"`
	GunReloadTime    *float64 `json:"gun_reload_time" db:"gun_reload_time"`
	GunTag           *string  `json:"gun_tag" db:"gun_tag"`
	GunTier          *int     `json:"gun_tier" db:"gun_tier"`
	GunTraverseSpeed *int     `json:"gun_traverse_speed" db:"gun_traverse_speed"`
	GunWeight        *int     `json:"gun_weight" db:"gun_weight"`

	// Mounted modules
	ModulesEngineID     *int `json:"modules_engine_id" db:"modules_engine_id"`
	ModulesGunID        *int `json:"modules_gun_id" db:"modules_gun_id"`
	ModulesRadioID      *int `json:"modules_radio_id" db:"modules_radio_id"`
	ModulesSuspensionID *int `json:"modules_suspension_id" db:"modules_suspension_id"`
	ModulesTurretID     *int `json:"modules_turret_id" db:"modules_turret_id"`

	// Radio characteristics
	RadioName        *string `json:"radio_name" db:"radio_name"`
	RadioSignalRange *int    `json:"radio_signal_range" db:"radio_signal_range"`
	RadioTag         *string `json:"radio_tag" db:"radio_tag"`
	RadioTier        *int    `json:"radio_tier" db:"radio_tier"`
	RadioWeight      *int    `json:"radio_weight" db:"radio_weight"`

	// Rapid mode (for wheeled vehicles) stored as JSONB
	Rapid JSONB `json:"rapid" db:"rapid"`

	// Siege mode characteristics stored as JSONB
	Siege JSONB `json:"siege" db:"siege"`

	// Suspension characteristics
	SuspensionLoadLimit         *int    `json:"suspension_load_limit" db:"suspension_load_limit"`
	SuspensionName              *string `json:"suspension_name" db:"suspension_name"`
	SuspensionSteeringLockAngle *int    `json:"suspension_steering_lock_angle" db:"suspension_steering_lock_angle"`
	SuspensionTag               *string `json:"suspension_tag" db:"suspension_tag"`
	SuspensionTier              *int    `json:"suspension_tier" db:"suspension_tier"`
	SuspensionTraverseSpeed     *int    `json:"suspension_traverse_speed" db:"suspension_traverse_speed"`
	SuspensionWeight            *int    `json:"suspension_weight" db:"suspension_weight"`

	// Turret characteristics
	TurretHP               *int    `json:"turret_hp" db:"turret_hp"`
	TurretName             *string `json:"turret_name" db:"turret_name"`
	TurretTag              *string `json:"turret_tag" db:"turret_tag"`
	TurretTier             *int    `json:"turret_tier" db:"turret_tier"`
	TurretTraverseLeftArc  *int    `json:"turret_traverse_left_arc" db:"turret_traverse_left_arc"`
	TurretTraverseRightArc *int    `json:"turret_traverse_right_arc" db:"turret_traverse_right_arc"`
	TurretTraverseSpeed    *int    `json:"turret_traverse_speed" db:"turret_traverse_speed"`
	TurretViewRange        *int    `json:"turret_view_range" db:"turret_view_range"`
	TurretWeight           *int    `json:"turret_weight" db:"turret_weight"`

	// Metadata
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Value implements the driver.Valuer interface for AmmoItem array
func (a AmmoItem) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan implements the sql.Scanner interface for AmmoItem
func (a *AmmoItem) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, a)
}

// ProfileFromDefaultProfile creates a Profile from API default_profile data
func ProfileFromDefaultProfile(tankID int, defaultProfile map[string]interface{}) *Profile {
	profile := &Profile{
		TankID: &tankID,
	}

	// Helper function to safely get int pointer
	getIntPtr := func(m map[string]interface{}, key string) *int {
		if val, ok := m[key]; ok {
			switch v := val.(type) {
			case float64:
				i := int(v)
				return &i
			case int:
				return &v
			}
		}
		return nil
	}

	// Helper function to safely get float64 pointer
	getFloat64Ptr := func(m map[string]interface{}, key string) *float64 {
		if val, ok := m[key]; ok {
			switch v := val.(type) {
			case float64:
				return &v
			case int:
				f := float64(v)
				return &f
			}
		}
		return nil
	}

	// Helper function to safely get string pointer
	getStringPtr := func(m map[string]interface{}, key string) *string {
		if val, ok := m[key]; ok {
			if s, ok := val.(string); ok {
				return &s
			}
		}
		return nil
	}

	// Basic characteristics
	profile.HP = getIntPtr(defaultProfile, "hp")
	profile.HullHP = getIntPtr(defaultProfile, "hull_hp")
	profile.HullWeight = getIntPtr(defaultProfile, "hull_weight")
	profile.MaxAmmo = getIntPtr(defaultProfile, "max_ammo")
	profile.MaxWeight = getIntPtr(defaultProfile, "max_weight")
	profile.SpeedBackward = getFloat64Ptr(defaultProfile, "speed_backward")
	profile.SpeedForward = getFloat64Ptr(defaultProfile, "speed_forward")
	profile.Weight = getIntPtr(defaultProfile, "weight")

	// Ammo
	if ammo, ok := defaultProfile["ammo"].([]interface{}); ok {
		ammoBytes, _ := json.Marshal(ammo)
		profile.Ammo = make(JSONB)
		json.Unmarshal(ammoBytes, &profile.Ammo)
	}

	// Armor
	if armor, ok := defaultProfile["armor"].(map[string]interface{}); ok {
		if hull, ok := armor["hull"].(map[string]interface{}); ok {
			profile.ArmorHullFront = getIntPtr(hull, "front")
			profile.ArmorHullRear = getIntPtr(hull, "rear")
			profile.ArmorHullSides = getIntPtr(hull, "sides")
		}
		if turret, ok := armor["turret"].(map[string]interface{}); ok {
			profile.ArmorTurretFront = getIntPtr(turret, "front")
			profile.ArmorTurretRear = getIntPtr(turret, "rear")
			profile.ArmorTurretSides = getIntPtr(turret, "sides")
		}
	}

	// Engine
	if engine, ok := defaultProfile["engine"].(map[string]interface{}); ok {
		profile.EngineFireChance = getFloat64Ptr(engine, "fire_chance")
		profile.EngineName = getStringPtr(engine, "name")
		profile.EnginePower = getIntPtr(engine, "power")
		profile.EngineTag = getStringPtr(engine, "tag")
		profile.EngineTier = getIntPtr(engine, "tier")
		profile.EngineWeight = getIntPtr(engine, "weight")
	}

	// Gun
	if gun, ok := defaultProfile["gun"].(map[string]interface{}); ok {
		profile.GunAimTime = getFloat64Ptr(gun, "aim_time")
		profile.GunCaliber = getIntPtr(gun, "caliber")
		profile.GunDispersion = getFloat64Ptr(gun, "dispersion")
		profile.GunFireRate = getFloat64Ptr(gun, "fire_rate")
		profile.GunMoveDownArc = getIntPtr(gun, "move_down_arc")
		profile.GunMoveUpArc = getIntPtr(gun, "move_up_arc")
		profile.GunName = getStringPtr(gun, "name")
		profile.GunReloadTime = getFloat64Ptr(gun, "reload_time")
		profile.GunTag = getStringPtr(gun, "tag")
		profile.GunTier = getIntPtr(gun, "tier")
		profile.GunTraverseSpeed = getIntPtr(gun, "traverse_speed")
		profile.GunWeight = getIntPtr(gun, "weight")
	}

	// Modules
	if modules, ok := defaultProfile["modules"].(map[string]interface{}); ok {
		profile.ModulesEngineID = getIntPtr(modules, "engine_id")
		profile.ModulesGunID = getIntPtr(modules, "gun_id")
		profile.ModulesRadioID = getIntPtr(modules, "radio_id")
		profile.ModulesSuspensionID = getIntPtr(modules, "suspension_id")
		profile.ModulesTurretID = getIntPtr(modules, "turret_id")
	}

	// Radio
	if radio, ok := defaultProfile["radio"].(map[string]interface{}); ok {
		profile.RadioName = getStringPtr(radio, "name")
		profile.RadioSignalRange = getIntPtr(radio, "signal_range")
		profile.RadioTag = getStringPtr(radio, "tag")
		profile.RadioTier = getIntPtr(radio, "tier")
		profile.RadioWeight = getIntPtr(radio, "weight")
	}

	// Rapid mode (for wheeled vehicles)
	if rapid, ok := defaultProfile["rapid"].(map[string]interface{}); ok {
		rapidBytes, _ := json.Marshal(rapid)
		profile.Rapid = make(JSONB)
		json.Unmarshal(rapidBytes, &profile.Rapid)
	}

	// Siege mode
	if siege, ok := defaultProfile["siege"].(map[string]interface{}); ok {
		siegeBytes, _ := json.Marshal(siege)
		profile.Siege = make(JSONB)
		json.Unmarshal(siegeBytes, &profile.Siege)
	}

	// Suspension
	if suspension, ok := defaultProfile["suspension"].(map[string]interface{}); ok {
		profile.SuspensionLoadLimit = getIntPtr(suspension, "load_limit")
		profile.SuspensionName = getStringPtr(suspension, "name")
		profile.SuspensionSteeringLockAngle = getIntPtr(suspension, "steering_lock_angle")
		profile.SuspensionTag = getStringPtr(suspension, "tag")
		profile.SuspensionTier = getIntPtr(suspension, "tier")
		profile.SuspensionTraverseSpeed = getIntPtr(suspension, "traverse_speed")
		profile.SuspensionWeight = getIntPtr(suspension, "weight")
	}

	// Turret
	if turret, ok := defaultProfile["turret"].(map[string]interface{}); ok {
		profile.TurretHP = getIntPtr(turret, "hp")
		profile.TurretName = getStringPtr(turret, "name")
		profile.TurretTag = getStringPtr(turret, "tag")
		profile.TurretTier = getIntPtr(turret, "tier")
		profile.TurretTraverseLeftArc = getIntPtr(turret, "traverse_left_arc")
		profile.TurretTraverseRightArc = getIntPtr(turret, "traverse_right_arc")
		profile.TurretTraverseSpeed = getIntPtr(turret, "traverse_speed")
		profile.TurretViewRange = getIntPtr(turret, "view_range")
		profile.TurretWeight = getIntPtr(turret, "weight")
	}

	return profile
}

// ModuleProfileFromDefaultProfile creates a Profile from Module API default_profile data
func ModuleProfileFromDefaultProfile(moduleID int, defaultProfile map[string]interface{}) *Profile {
	profile := &Profile{
		ModuleID: &moduleID,
	}

	// Helper functions (same as above)
	getIntPtr := func(m map[string]interface{}, key string) *int {
		if val, ok := m[key]; ok {
			switch v := val.(type) {
			case float64:
				i := int(v)
				return &i
			case int:
				return &v
			}
		}
		return nil
	}

	getFloat64Ptr := func(m map[string]interface{}, key string) *float64 {
		if val, ok := m[key]; ok {
			switch v := val.(type) {
			case float64:
				return &v
			case int:
				f := float64(v)
				return &f
			}
		}
		return nil
	}

	// Engine characteristics
	if engine, ok := defaultProfile["engine"].(map[string]interface{}); ok {
		profile.EngineFireChance = getFloat64Ptr(engine, "fire_chance")
		profile.EnginePower = getIntPtr(engine, "power")
	}

	// Gun characteristics
	if gun, ok := defaultProfile["gun"].(map[string]interface{}); ok {
		profile.GunAimTime = getFloat64Ptr(gun, "aim_time")
		profile.GunDispersion = getFloat64Ptr(gun, "dispersion")
		profile.GunFireRate = getFloat64Ptr(gun, "fire_rate")
		profile.GunMoveDownArc = getIntPtr(gun, "move_down_arc")
		profile.GunMoveUpArc = getIntPtr(gun, "move_up_arc")
		profile.GunReloadTime = getFloat64Ptr(gun, "reload_time")
		profile.GunTraverseSpeed = getIntPtr(gun, "traverse_speed")
		profile.MaxAmmo = getIntPtr(gun, "max_ammo")

		// Gun ammo
		if ammo, ok := gun["ammo"].([]interface{}); ok {
			ammoBytes, _ := json.Marshal(ammo)
			profile.Ammo = make(JSONB)
			json.Unmarshal(ammoBytes, &profile.Ammo)
		}
	}

	// Radio characteristics
	if radio, ok := defaultProfile["radio"].(map[string]interface{}); ok {
		profile.RadioSignalRange = getIntPtr(radio, "signal_range")
	}

	// Suspension characteristics
	if suspension, ok := defaultProfile["suspension"].(map[string]interface{}); ok {
		profile.SuspensionLoadLimit = getIntPtr(suspension, "load_limit")
		profile.SuspensionTraverseSpeed = getIntPtr(suspension, "traverse_speed")
	}

	// Turret characteristics
	if turret, ok := defaultProfile["turret"].(map[string]interface{}); ok {
		profile.TurretHP = getIntPtr(turret, "hp")
		profile.TurretTraverseSpeed = getIntPtr(turret, "traverse_speed")
		profile.TurretViewRange = getIntPtr(turret, "view_range")
		profile.ArmorTurretFront = getIntPtr(turret, "armor_front")
		profile.ArmorTurretRear = getIntPtr(turret, "armor_rear")
		profile.ArmorTurretSides = getIntPtr(turret, "armor_sides")
	}

	return profile
}

// ProfileRepository defines data access operations for profiles
type ProfileRepository interface {
	// GetProfileByTankID retrieves a profile by tank_id
	GetProfileByTankID(tankID int) (*Profile, error)

	// CreateProfile creates a new profile
	CreateProfile(profile *Profile) error

	// UpdateProfile updates an existing profile
	UpdateProfile(profile *Profile) error

	// DeleteProfile deletes a profile by tank_id
	DeleteProfile(tankID int) error

	// GetAllProfiles retrieves all profiles
	GetAllProfiles() ([]Profile, error)
}

// ProfileUseCase defines business logic operations for profiles
type ProfileUseCase interface {
	// GetProfileByTankID retrieves a profile by tank_id
	GetProfileByTankID(tankID int) (*Profile, error)

	// CreateProfile creates a new profile
	CreateProfile(profile *Profile) error

	// UpdateProfile updates an existing profile
	UpdateProfile(profile *Profile) error

	// DeleteProfile deletes a profile by tank_id
	DeleteProfile(tankID int) error

	// GetAllProfiles retrieves all profiles
	GetAllProfiles() ([]Profile, error)
}
