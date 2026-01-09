package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"wot-statistics-server/domain"
	"wot-statistics-server/internal/infrastructure/postgres"
)

type profileRepository struct {
	db *postgres.DB
}

// NewProfileRepository creates a new profile repository instance
func NewProfileRepository(db *postgres.DB) domain.ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

// GetProfileByTankID retrieves a profile by tank_id
func (r *profileRepository) GetProfileByTankID(tankID int) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT 
			id, tank_id, module_id, hp, hull_hp, hull_weight, max_ammo, max_weight,
			speed_backward, speed_forward, weight, ammo,
			armor_hull_front, armor_hull_rear, armor_hull_sides,
			armor_turret_front, armor_turret_rear, armor_turret_sides,
			engine_fire_chance, engine_name, engine_power, engine_tag, engine_tier, engine_weight,
			gun_aim_time, gun_caliber, gun_dispersion, gun_fire_rate,
			gun_move_down_arc, gun_move_up_arc, gun_name, gun_reload_time,
			gun_tag, gun_tier, gun_traverse_speed, gun_weight,
			modules_engine_id, modules_gun_id, modules_radio_id,
			modules_suspension_id, modules_turret_id,
			radio_name, radio_signal_range, radio_tag, radio_tier, radio_weight,
			rapid, siege,
			suspension_load_limit, suspension_name, suspension_steering_lock_angle,
			suspension_tag, suspension_tier, suspension_traverse_speed, suspension_weight,
			turret_hp, turret_name, turret_tag, turret_tier,
			turret_traverse_left_arc, turret_traverse_right_arc,
			turret_traverse_speed, turret_view_range, turret_weight,
			created_at, updated_at
		FROM profiles
		WHERE tank_id = $1
	`

	var profile domain.Profile
	err := r.db.GetContext(ctx, &profile, query, tankID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("profile for tank ID %d not found", tankID)
		}
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	return &profile, nil
}

// GetProfileByModuleID retrieves a profile by module_id
func (r *profileRepository) GetProfileByModuleID(moduleID int) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT 
			id, tank_id, module_id, hp, hull_hp, hull_weight, max_ammo, max_weight,
			speed_backward, speed_forward, weight, ammo,
			armor_hull_front, armor_hull_rear, armor_hull_sides,
			armor_turret_front, armor_turret_rear, armor_turret_sides,
			engine_fire_chance, engine_name, engine_power, engine_tag, engine_tier, engine_weight,
			gun_aim_time, gun_caliber, gun_dispersion, gun_fire_rate,
			gun_move_down_arc, gun_move_up_arc, gun_name, gun_reload_time,
			gun_tag, gun_tier, gun_traverse_speed, gun_weight,
			modules_engine_id, modules_gun_id, modules_radio_id,
			modules_suspension_id, modules_turret_id,
			radio_name, radio_signal_range, radio_tag, radio_tier, radio_weight,
			rapid, siege,
			suspension_load_limit, suspension_name, suspension_steering_lock_angle,
			suspension_tag, suspension_tier, suspension_traverse_speed, suspension_weight,
			turret_hp, turret_name, turret_tag, turret_tier,
			turret_traverse_left_arc, turret_traverse_right_arc,
			turret_traverse_speed, turret_view_range, turret_weight,
			created_at, updated_at
		FROM profiles
		WHERE module_id = $1
	`

	var profile domain.Profile
	err := r.db.GetContext(ctx, &profile, query, moduleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("profile for module ID %d not found", moduleID)
		}
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	return &profile, nil
}

// CreateProfile creates a new profile
func (r *profileRepository) CreateProfile(profile *domain.Profile) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO profiles (
			tank_id, module_id, hp, hull_hp, hull_weight, max_ammo, max_weight,
			speed_backward, speed_forward, weight, ammo,
			armor_hull_front, armor_hull_rear, armor_hull_sides,
			armor_turret_front, armor_turret_rear, armor_turret_sides,
			engine_fire_chance, engine_name, engine_power, engine_tag, engine_tier, engine_weight,
			gun_aim_time, gun_caliber, gun_dispersion, gun_fire_rate,
			gun_move_down_arc, gun_move_up_arc, gun_name, gun_reload_time,
			gun_tag, gun_tier, gun_traverse_speed, gun_weight,
			modules_engine_id, modules_gun_id, modules_radio_id,
			modules_suspension_id, modules_turret_id,
			radio_name, radio_signal_range, radio_tag, radio_tier, radio_weight,
			rapid, siege,
			suspension_load_limit, suspension_name, suspension_steering_lock_angle,
			suspension_tag, suspension_tier, suspension_traverse_speed, suspension_weight,
			turret_hp, turret_name, turret_tag, turret_tier,
			turret_traverse_left_arc, turret_traverse_right_arc,
			turret_traverse_speed, turret_view_range, turret_weight
		) VALUES (
			:tank_id, :module_id, :hp, :hull_hp, :hull_weight, :max_ammo, :max_weight,
			:speed_backward, :speed_forward, :weight, :ammo,
			:armor_hull_front, :armor_hull_rear, :armor_hull_sides,
			:armor_turret_front, :armor_turret_rear, :armor_turret_sides,
			:engine_fire_chance, :engine_name, :engine_power, :engine_tag, :engine_tier, :engine_weight,
			:gun_aim_time, :gun_caliber, :gun_dispersion, :gun_fire_rate,
			:gun_move_down_arc, :gun_move_up_arc, :gun_name, :gun_reload_time,
			:gun_tag, :gun_tier, :gun_traverse_speed, :gun_weight,
			:modules_engine_id, :modules_gun_id, :modules_radio_id,
			:modules_suspension_id, :modules_turret_id,
			:radio_name, :radio_signal_range, :radio_tag, :radio_tier, :radio_weight,
			:rapid, :siege,
			:suspension_load_limit, :suspension_name, :suspension_steering_lock_angle,
			:suspension_tag, :suspension_tier, :suspension_traverse_speed, :suspension_weight,
			:turret_hp, :turret_name, :turret_tag, :turret_tier,
			:turret_traverse_left_arc, :turret_traverse_right_arc,
			:turret_traverse_speed, :turret_view_range, :turret_weight
		) RETURNING id
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, &profile.ID, profile)
	if err != nil {
		return fmt.Errorf("failed to create profile: %w", err)
	}

	return nil
}

// UpdateProfile updates an existing profile
func (r *profileRepository) UpdateProfile(profile *domain.Profile) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE profiles SET
			hp = :hp, hull_hp = :hull_hp, hull_weight = :hull_weight,
			max_ammo = :max_ammo, max_weight = :max_weight,
			speed_backward = :speed_backward, speed_forward = :speed_forward,
			weight = :weight, ammo = :ammo,
			armor_hull_front = :armor_hull_front, armor_hull_rear = :armor_hull_rear,
			armor_hull_sides = :armor_hull_sides,
			armor_turret_front = :armor_turret_front, armor_turret_rear = :armor_turret_rear,
			armor_turret_sides = :armor_turret_sides,
			engine_fire_chance = :engine_fire_chance, engine_name = :engine_name,
			engine_power = :engine_power, engine_tag = :engine_tag,
			engine_tier = :engine_tier, engine_weight = :engine_weight,
			gun_aim_time = :gun_aim_time, gun_caliber = :gun_caliber,
			gun_dispersion = :gun_dispersion, gun_fire_rate = :gun_fire_rate,
			gun_move_down_arc = :gun_move_down_arc, gun_move_up_arc = :gun_move_up_arc,
			gun_name = :gun_name, gun_reload_time = :gun_reload_time,
			gun_tag = :gun_tag, gun_tier = :gun_tier,
			gun_traverse_speed = :gun_traverse_speed, gun_weight = :gun_weight,
			modules_engine_id = :modules_engine_id, modules_gun_id = :modules_gun_id,
			modules_radio_id = :modules_radio_id, modules_suspension_id = :modules_suspension_id,
			modules_turret_id = :modules_turret_id,
			radio_name = :radio_name, radio_signal_range = :radio_signal_range,
			radio_tag = :radio_tag, radio_tier = :radio_tier, radio_weight = :radio_weight,
			rapid = :rapid, siege = :siege,
			suspension_load_limit = :suspension_load_limit, suspension_name = :suspension_name,
			suspension_steering_lock_angle = :suspension_steering_lock_angle,
			suspension_tag = :suspension_tag, suspension_tier = :suspension_tier,
			suspension_traverse_speed = :suspension_traverse_speed,
			suspension_weight = :suspension_weight,
			turret_hp = :turret_hp, turret_name = :turret_name,
			turret_tag = :turret_tag, turret_tier = :turret_tier,
			turret_traverse_left_arc = :turret_traverse_left_arc,
			turret_traverse_right_arc = :turret_traverse_right_arc,
			turret_traverse_speed = :turret_traverse_speed,
			turret_view_range = :turret_view_range, turret_weight = :turret_weight,
			updated_at = CURRENT_TIMESTAMP
		WHERE tank_id = :tank_id
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, profile)
	if err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("profile with tank_id %d not found", profile.TankID)
	}

	return nil
}

// DeleteProfile deletes a profile by tank_id
func (r *profileRepository) DeleteProfile(tankID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM profiles WHERE tank_id = $1`

	result, err := r.db.ExecContext(ctx, query, tankID)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("profile with tank_id %d not found", tankID)
	}

	return nil
}

// GetAllProfiles retrieves all profiles
func (r *profileRepository) GetAllProfiles() ([]domain.Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT 
			id, tank_id, hp, hull_hp, hull_weight, max_ammo, max_weight,
			speed_backward, speed_forward, weight, ammo,
			armor_hull_front, armor_hull_rear, armor_hull_sides,
			armor_turret_front, armor_turret_rear, armor_turret_sides,
			engine_fire_chance, engine_name, engine_power, engine_tag, engine_tier, engine_weight,
			gun_aim_time, gun_caliber, gun_dispersion, gun_fire_rate,
			gun_move_down_arc, gun_move_up_arc, gun_name, gun_reload_time,
			gun_tag, gun_tier, gun_traverse_speed, gun_weight,
			modules_engine_id, modules_gun_id, modules_radio_id,
			modules_suspension_id, modules_turret_id,
			radio_name, radio_signal_range, radio_tag, radio_tier, radio_weight,
			rapid, siege,
			suspension_load_limit, suspension_name, suspension_steering_lock_angle,
			suspension_tag, suspension_tier, suspension_traverse_speed, suspension_weight,
			turret_hp, turret_name, turret_tag, turret_tier,
			turret_traverse_left_arc, turret_traverse_right_arc,
			turret_traverse_speed, turret_view_range, turret_weight,
			created_at, updated_at
		FROM profiles
		ORDER BY tank_id
	`

	var profiles []domain.Profile
	err := r.db.SelectContext(ctx, &profiles, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get profiles: %w", err)
	}

	return profiles, nil
}
