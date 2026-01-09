package repository

import (
	"context"
	"fmt"
	"time"
	"wot-statistics-server/domain"
	"wot-statistics-server/internal/infrastructure/postgres"
)

type moduleRepository struct {
	db *postgres.DB
}

// NewModuleRepository creates a new instance of ModuleRepository
func NewModuleRepository(db *postgres.DB) domain.ModuleRepository {
	return &moduleRepository{
		db: db,
	}
}

// GetModules retrieves all modules from database
func (r *moduleRepository) GetModules() ([]domain.Module, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, module_id, name, type, nation, tier, weight, price_credit, image, created_at, updated_at
		FROM modules
		ORDER BY tier, type, nation, name
	`

	var modules []domain.Module
	err := r.db.SelectContext(ctx, &modules, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get modules: %w", err)
	}

	return modules, nil
}

// GetModuleByID retrieves a single module by its module_id
func (r *moduleRepository) GetModuleByID(moduleID int) (*domain.Module, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT id, module_id, name, type, nation, tier, weight, price_credit, image, created_at, updated_at
		FROM modules
		WHERE module_id = $1
	`

	var module domain.Module
	err := r.db.GetContext(ctx, &module, query, moduleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get module by ID %d: %w", moduleID, err)
	}

	return &module, nil
}

// GetModulesByFilter retrieves modules with filters
func (r *moduleRepository) GetModulesByFilter(moduleType string, nation string, tier int) ([]domain.Module, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, module_id, name, type, nation, tier, weight, price_credit, image, created_at, updated_at
		FROM modules
		WHERE 1=1
	`
	args := []interface{}{}
	argIndex := 1

	if moduleType != "" {
		query += fmt.Sprintf(" AND type = $%d", argIndex)
		args = append(args, moduleType)
		argIndex++
	}

	if nation != "" {
		query += fmt.Sprintf(" AND nation = $%d", argIndex)
		args = append(args, nation)
		argIndex++
	}

	if tier > 0 {
		query += fmt.Sprintf(" AND tier = $%d", argIndex)
		args = append(args, tier)
		argIndex++
	}

	query += " ORDER BY tier, type, nation, name"

	var modules []domain.Module
	err := r.db.SelectContext(ctx, &modules, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get modules by filter: %w", err)
	}

	return modules, nil
}

// CreateModules inserts multiple modules into database (bulk insert with UPSERT)
func (r *moduleRepository) CreateModules(modules []*domain.Module) (int, error) {
	if len(modules) == 0 {
		return 0, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Start transaction
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
		INSERT INTO modules (module_id, name, type, nation, tier, weight, price_credit, image)
		VALUES (:module_id, :name, :type, :nation, :tier, :weight, :price_credit, :image)
		ON CONFLICT (module_id) DO UPDATE SET
			name = EXCLUDED.name,
			type = EXCLUDED.type,
			nation = EXCLUDED.nation,
			tier = EXCLUDED.tier,
			weight = EXCLUDED.weight,
			price_credit = EXCLUDED.price_credit,
			image = EXCLUDED.image,
			updated_at = CURRENT_TIMESTAMP
	`

	inserted := 0
	for _, module := range modules {
		_, err := tx.NamedExecContext(ctx, query, module)
		if err != nil {
			return inserted, fmt.Errorf("failed to insert module %d: %w", module.ModuleID, err)
		}
		inserted++
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return inserted, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return inserted, nil
}

// ModuleExists checks if a module exists by module_id
func (r *moduleRepository) ModuleExists(moduleID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT EXISTS(SELECT 1 FROM modules WHERE module_id = $1)`

	var exists bool
	err := r.db.GetContext(ctx, &exists, query, moduleID)
	if err != nil {
		return false, fmt.Errorf("failed to check module existence: %w", err)
	}

	return exists, nil
}
