# Database Migrations

This directory contains SQL migration files for the WoT Statistics Server database.

## Migration Files

### 001_create_tanks_table
- **Up**: Creates the `tanks` table with all necessary columns and indexes
- **Down**: Drops the `tanks` table and related objects

## Running Migrations

### Using golang-migrate CLI

1. Install golang-migrate:
```bash
# macOS
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

# Or using Go
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

2. Run migrations:
```bash
# Up (apply migrations)
migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/wot-statistics?sslmode=disable" up

# Down (rollback migrations)
migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/wot-statistics?sslmode=disable" down

# Specific version
migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/wot-statistics?sslmode=disable" goto 1

# Force version (if stuck)
migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/wot-statistics?sslmode=disable" force 1
```

### Using psql directly

```bash
# Run up migration
psql -U postgres -d wot-statistics -f migrations/001_create_tanks_table.up.sql

# Run down migration
psql -U postgres -d wot-statistics -f migrations/001_create_tanks_table.down.sql
```

### Using Docker

```bash
# Run migrations in Docker container
docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/001_create_tanks_table.up.sql
```

## Schema Overview

### tanks table
Stores World of Tanks vehicle encyclopedia data from the Wargaming API.

**Key columns:**
- `tank_id`: Unique vehicle ID from API
- `name`, `short_name`: Vehicle names
- `nation`: Vehicle nation (ussr, germany, usa, etc.)
- `tier`: Vehicle tier (1-10)
- `type`: Vehicle type (heavyTank, mediumTank, etc.)
- `is_premium`, `is_gift`: Premium status
- `price_credit`, `price_gold`: Purchase prices
- `default_profile`: JSONB with technical specs
- Various module arrays stored as JSONB

**Indexes:**
- Single column indexes on: tank_id, nation, tier, type, is_premium, name
- Composite indexes for common filter combinations
- GIN indexes on JSONB columns for efficient querying
