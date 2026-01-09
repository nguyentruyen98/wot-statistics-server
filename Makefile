# Variables
MAIN_PATH=./cmd/api
SEED_PATH=./cmd/seed

# Color output
CYAN=\033[0;36m
GREEN=\033[0;32m
YELLOW=\033[1;33m
NC=\033[0m # No Color

## run: Run the application locally
run:
	@echo "$(CYAN)Running application...$(NC)"
	@go run $(MAIN_PATH)/main.go

## seed: Seed the database with tank data (full setup)
seed:
	@echo "$(GREEN)Seeding database with tank data...$(NC)"
	@chmod +x scripts/seed.sh
	@./scripts/seed.sh

## seed-quick: Quick seed (assumes DB is already set up)
seed-quick:
	@echo "$(YELLOW)Quick seeding tank data...$(NC)"
	@chmod +x scripts/seed-quick.sh
	@./scripts/seed-quick.sh

## seed-go: Run seed directly with go run
seed-go:
	@echo "$(CYAN)Running seeder...$(NC)"
	@DATABASE_URL="postgres://postgres:postgres@localhost:5432/wot-statistics?sslmode=disable" \
	 TANKS_JSON_FILE="tanks_X.json" \
	 go run $(SEED_PATH)/main.go

## db-setup: Setup database (create + migrate)
db-setup:
	@echo "$(GREEN)Setting up database...$(NC)"
	@docker-compose up -d postgres
	@sleep 3
	@docker exec wot-statistics-postgres psql -U postgres -c "CREATE DATABASE \"wot-statistics\";" 2>/dev/null || echo "Database already exists"
	@echo "$(CYAN)Running migrations...$(NC)"
	@docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/001_create_tanks_table.up.sql 2>&1 | grep -v "already exists" || true
	@docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/002_create_modules_table.up.sql 2>&1 | grep -v "already exists" || true
	@docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/003_create_profiles_table.up.sql 2>&1 | grep -v "already exists" || true
	@docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/004_drop_default_profile_from_tanks.up.sql 2>&1 | grep -v "does not exist" || true
	@docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/005_add_module_support_to_profiles.up.sql 2>&1 | grep -v "already exists" || true
	@echo "$(GREEN)Database setup completed!$(NC)"

## db-reset: Reset database (drop, create, migrate)
db-reset:
	@echo "$(YELLOW)Resetting database...$(NC)"
	@docker exec wot-statistics-postgres psql -U postgres -c "DROP DATABASE IF EXISTS \"wot-statistics\";"
	@docker exec wot-statistics-postgres psql -U postgres -c "CREATE DATABASE \"wot-statistics\";"
	@echo "$(CYAN)Running migrations...$(NC)"
	@docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/001_create_tanks_table.up.sql
	@docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/002_create_modules_table.up.sql
	@docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/003_create_profiles_table.up.sql
	@echo "$(GREEN)Database reset completed!$(NC)"

## migrate-up: Run all pending migrations on existing database
migrate-up:
	@echo "$(CYAN)Running pending migrations...$(NC)"
	@docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/001_create_tanks_table.up.sql 2>&1 | grep -v "already exists" || true
	@docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/002_create_modules_table.up.sql 2>&1 | grep -v "already exists" || true
	@docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/003_create_profiles_table.up.sql 2>&1 | grep -v "already exists" || true
	@docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/004_drop_default_profile_from_tanks.up.sql 2>&1 | grep -v "does not exist" || true
	@docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/005_add_module_support_to_profiles.up.sql 2>&1 | grep -v "already exists" || true
	@echo "$(GREEN)Migrations completed!$(NC)"

## db-psql: Open PostgreSQL shell
db-psql:
	@docker exec -it wot-statistics-postgres psql -U postgres -d wot-statistics

## help: Display this help message
help:
	@echo "Available commands:"
	@echo "  make run          - Run the application"
	@echo "  make seed         - Seed database (full setup with checks)"
	@echo "  make seed-quick   - Quick seed (DB must be ready)"
	@echo "  make seed-go      - Run seeder directly with go run"
	@echo "  make db-setup     - Setup database (create + migrate)"
	@echo "  make db-reset     - Reset database (drop + create + migrate)"
	@echo "  make migrate-up   - Run all pending migrations"
	@echo "  make db-psql      - Open PostgreSQL shell"
	@echo "  make help         - Show this help message"

.PHONY: run seed seed-quick seed-go db-setup db-reset migrate-up db-psql help