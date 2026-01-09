#!/bin/bash

# Tank Data Seeder Script
# This script seeds the database with tank data from tanks_X.json

set -e  # Exit on error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  WoT Statistics - Tank Data Seeder${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""

# Check if Docker container is running
echo -e "${YELLOW}Checking PostgreSQL container...${NC}"
if ! docker ps | grep -q wot-statistics-postgres; then
    echo -e "${RED}PostgreSQL container is not running!${NC}"
    echo -e "${YELLOW}Starting PostgreSQL with docker-compose...${NC}"
    docker-compose up -d postgres
    echo -e "${GREEN}Waiting for PostgreSQL to be ready...${NC}"
    sleep 5
fi

# Check if database exists
echo -e "${YELLOW}Checking if database exists...${NC}"
DB_EXISTS=$(docker exec wot-statistics-postgres psql -U postgres -tAc "SELECT 1 FROM pg_database WHERE datname='wot-statistics'" 2>/dev/null || echo "0")

if [ "$DB_EXISTS" != "1" ]; then
    echo -e "${YELLOW}Database does not exist. Creating...${NC}"
    docker exec wot-statistics-postgres psql -U postgres -c "CREATE DATABASE \"wot-statistics\";"
    echo -e "${GREEN}Database created!${NC}"
else
    echo -e "${GREEN}Database already exists.${NC}"
fi

# Run migrations
echo -e "${YELLOW}Running database migrations...${NC}"
if [ -f "migrations/001_create_tanks_table.up.sql" ]; then
    docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/001_create_tanks_table.up.sql 2>&1 | grep -v "already exists" || true
    echo -e "${GREEN}Migrations completed!${NC}"
else
    echo -e "${RED}Migration file not found!${NC}"
    exit 1
fi

# Check if tanks_X.json exists
if [ ! -f "tanks_X.json" ]; then
    echo -e "${RED}Error: tanks_X.json file not found!${NC}"
    exit 1
fi

# Set environment variables
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/wot-statistics?sslmode=disable"
export TANKS_JSON_FILE="tanks_X.json"

# Run the seeder
echo -e "${YELLOW}Running tank data seeder...${NC}"
echo ""
go run cmd/seed/main.go

echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  Seeding completed successfully!${NC}"
echo -e "${GREEN}========================================${NC}"
