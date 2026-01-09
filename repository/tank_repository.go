package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"wot-statistics-server/config"
	"wot-statistics-server/domain"
	"wot-statistics-server/external"
	"wot-statistics-server/internal/infrastructure/postgres"
)

type tankRepository struct {
	db           *postgres.DB
	wargamingAPI *external.WargamingAPI
}

// NewTankRepository creates a new tank repository instance
func NewTankRepository(db *postgres.DB, appConfig *config.AppConfig) domain.TankRepository {
	return &tankRepository{
		db:           db,
		wargamingAPI: external.NewWargamingAPI(appConfig),
	}
}

// GetTanks retrieves all tanks from database
func (r *tankRepository) GetTanks() ([]domain.Tank, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT 
			id, tank_id, name, short_name, nation, tier, type, tag,
			is_premium, is_premium_igr, is_gift, is_wheeled,
			price_credit, price_gold, description, big_icon_url,
			radios, suspensions, provisions, engines, guns, turrets, crew,
			modules_tree, next_tanks, prices_xp,
			multination, created_at, updated_at
		FROM tanks
		ORDER BY nation, tier, name
	`

	var tanks []domain.Tank
	err := r.db.SelectContext(ctx, &tanks, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get tanks: %w", err)
	}

	return tanks, nil
}

// GetTankByID retrieves a single tank by its tank_id
func (r *tankRepository) GetTankByID(tankID int) (*domain.Tank, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT 
			id, tank_id, name, short_name, nation, tier, type, tag,
			is_premium, is_premium_igr, is_gift, is_wheeled,
			price_credit, price_gold, description, big_icon_url,
			radios, suspensions, provisions, engines, guns, turrets, crew,
			modules_tree, next_tanks, prices_xp,
			multination, created_at, updated_at
		FROM tanks
		WHERE tank_id = $1
	`

	var tank domain.Tank
	err := r.db.GetContext(ctx, &tank, query, tankID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tank with ID %d not found", tankID)
		}
		return nil, fmt.Errorf("failed to get tank: %w", err)
	}

	return &tank, nil
}

// GetTanksByFilter retrieves tanks with optional filters
func (r *tankRepository) GetTanksByFilter(nation string, tier int, tankType string, isPremium *bool) ([]domain.Tank, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Build dynamic query with filters
	conditions := []string{}
	args := []interface{}{}
	argIndex := 1

	if nation != "" {
		conditions = append(conditions, fmt.Sprintf("nation = $%d", argIndex))
		args = append(args, nation)
		argIndex++
	}

	if tier > 0 {
		conditions = append(conditions, fmt.Sprintf("tier = $%d", argIndex))
		args = append(args, tier)
		argIndex++
	}

	if tankType != "" {
		conditions = append(conditions, fmt.Sprintf("type = $%d", argIndex))
		args = append(args, tankType)
		argIndex++
	}

	if isPremium != nil {
		conditions = append(conditions, fmt.Sprintf("is_premium = $%d", argIndex))
		args = append(args, *isPremium)
		argIndex++
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	query := fmt.Sprintf(`
		SELECT 
			id, tank_id, name, short_name, nation, tier, type, tag,
			is_premium, is_premium_igr, is_gift, is_wheeled,
			price_credit, price_gold, description, big_icon_url,
			radios, suspensions, provisions, engines, guns, turrets, crew,
			modules_tree, next_tanks, prices_xp,
			multination, created_at, updated_at
		FROM tanks
		%s
		ORDER BY tier, nation, name
	`, whereClause)

	var tanks []domain.Tank
	err := r.db.SelectContext(ctx, &tanks, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get tanks by filter: %w", err)
	}

	return tanks, nil
}

// CreateTank inserts a new tank into database
func (r *tankRepository) CreateTank(tank *domain.Tank) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO tanks (
			tank_id, name, short_name, nation, tier, type, tag,
			is_premium, is_premium_igr, is_gift, is_wheeled,
			price_credit, price_gold, description, big_icon_url,
			radios, suspensions, provisions, engines, guns, turrets, crew,
			default_profile, modules_tree, next_tanks, prices_xp,
			multination
		) VALUES (
			:tank_id, :name, :short_name, :nation, :tier, :type, :tag,
			:is_premium, :is_premium_igr, :is_gift, :is_wheeled,
			:price_credit, :price_gold, :description, :big_icon_url,
			:radios, :suspensions, :provisions, :engines, :guns, :turrets, :crew,
			:default_profile, :modules_tree, :next_tanks, :prices_xp,
			:multination
		)
		ON CONFLICT (tank_id) DO UPDATE SET
			name = EXCLUDED.name,
			short_name = EXCLUDED.short_name,
			nation = EXCLUDED.nation,
			tier = EXCLUDED.tier,
			type = EXCLUDED.type,
			tag = EXCLUDED.tag,
			is_premium = EXCLUDED.is_premium,
			is_premium_igr = EXCLUDED.is_premium_igr,
			is_gift = EXCLUDED.is_gift,
			is_wheeled = EXCLUDED.is_wheeled,
			price_credit = EXCLUDED.price_credit,
			price_gold = EXCLUDED.price_gold,
			description = EXCLUDED.description,
			big_icon_url = EXCLUDED.big_icon_url,
			radios = EXCLUDED.radios,
			suspensions = EXCLUDED.suspensions,
			provisions = EXCLUDED.provisions,
			engines = EXCLUDED.engines,
			guns = EXCLUDED.guns,
			turrets = EXCLUDED.turrets,
			crew = EXCLUDED.crew,
			modules_tree = EXCLUDED.modules_tree,
			next_tanks = EXCLUDED.next_tanks,
			prices_xp = EXCLUDED.prices_xp,
			multination = EXCLUDED.multination,
			updated_at = CURRENT_TIMESTAMP
		RETURNING id
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, &tank.ID, tank)
	if err != nil {
		return fmt.Errorf("failed to create tank: %w", err)
	}

	return nil
}

// CreateTanks inserts multiple tanks into database (bulk insert with transaction)
func (r *tankRepository) CreateTanks(tanks []*domain.Tank) (int, error) {
	if len(tanks) == 0 {
		return 0, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Start transaction
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
		INSERT INTO tanks (
			tank_id, name, short_name, nation, tier, type, tag,
			is_premium, is_premium_igr, is_gift, is_wheeled,
			price_credit, price_gold, description, big_icon_url,
			radios, suspensions, provisions, engines, guns, turrets, crew,
			modules_tree, next_tanks, prices_xp,
			multination
		) VALUES (
			:tank_id, :name, :short_name, :nation, :tier, :type, :tag,
			:is_premium, :is_premium_igr, :is_gift, :is_wheeled,
			:price_credit, :price_gold, :description, :big_icon_url,
			:radios, :suspensions, :provisions, :engines, :guns, :turrets, :crew,
			:modules_tree, :next_tanks, :prices_xp,
			:multination
		)
		ON CONFLICT (tank_id) DO UPDATE SET
			name = EXCLUDED.name,
			short_name = EXCLUDED.short_name,
			nation = EXCLUDED.nation,
			tier = EXCLUDED.tier,
			type = EXCLUDED.type,
			tag = EXCLUDED.tag,
			is_premium = EXCLUDED.is_premium,
			is_premium_igr = EXCLUDED.is_premium_igr,
			is_gift = EXCLUDED.is_gift,
			is_wheeled = EXCLUDED.is_wheeled,
			price_credit = EXCLUDED.price_credit,
			price_gold = EXCLUDED.price_gold,
			description = EXCLUDED.description,
			big_icon_url = EXCLUDED.big_icon_url,
			radios = EXCLUDED.radios,
			suspensions = EXCLUDED.suspensions,
			provisions = EXCLUDED.provisions,
			engines = EXCLUDED.engines,
			guns = EXCLUDED.guns,
			turrets = EXCLUDED.turrets,
			crew = EXCLUDED.crew,
			modules_tree = EXCLUDED.modules_tree,
			next_tanks = EXCLUDED.next_tanks,
			prices_xp = EXCLUDED.prices_xp,
			multination = EXCLUDED.multination,
			updated_at = CURRENT_TIMESTAMP
	`

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	insertedCount := 0
	for _, tank := range tanks {
		_, err := stmt.ExecContext(ctx, tank)
		if err != nil {
			return insertedCount, fmt.Errorf("failed to insert tank %d: %w", tank.TankID, err)
		}
		insertedCount++
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return insertedCount, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return insertedCount, nil
}

// UpdateTank updates an existing tank in database
func (r *tankRepository) UpdateTank(tank *domain.Tank) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE tanks SET
			name = :name,
			short_name = :short_name,
			nation = :nation,
			tier = :tier,
			type = :type,
			tag = :tag,
			is_premium = :is_premium,
			is_premium_igr = :is_premium_igr,
			is_gift = :is_gift,
			is_wheeled = :is_wheeled,
			price_credit = :price_credit,
			price_gold = :price_gold,
			description = :description,
			big_icon_url = :big_icon_url,
			radios = :radios,
			suspensions = :suspensions,
			provisions = :provisions,
			engines = :engines,
			guns = :guns,
			turrets = :turrets,
			crew = :crew,
			modules_tree = :modules_tree,
			next_tanks = :next_tanks,
			prices_xp = :prices_xp,
			multination = :multination,
			updated_at = CURRENT_TIMESTAMP
		WHERE tank_id = :tank_id
	`

	result, err := r.db.NamedExecContext(ctx, query, tank)
	if err != nil {
		return fmt.Errorf("failed to update tank: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("tank with ID %d not found", tank.TankID)
	}

	return nil
}

// DeleteTank deletes a tank by tank_id
func (r *tankRepository) DeleteTank(tankID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM tanks WHERE tank_id = $1`

	result, err := r.db.ExecContext(ctx, query, tankID)
	if err != nil {
		return fmt.Errorf("failed to delete tank: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("tank with ID %d not found", tankID)
	}

	return nil
}

// TankExists checks if a tank exists by tank_id
func (r *tankRepository) TankExists(tankID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM tanks WHERE tank_id = $1)`

	err := r.db.GetContext(ctx, &exists, query, tankID)
	if err != nil {
		return false, fmt.Errorf("failed to check tank existence: %w", err)
	}

	return exists, nil
}

// GetTanksFromAPI fetches tanks from Wargaming API
func (r *tankRepository) GetTanksFromAPI() (*external.APIResponse, error) {
	params := map[string]string{
		"limit":  "100",
		"fields": "tank_id,name,short_name,type,tier,nation,is_premium,description,images",
	}
	return r.wargamingAPI.GetTanksList(params)
}
