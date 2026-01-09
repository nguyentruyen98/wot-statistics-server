#!/bin/bash

# Quick seed script (assumes database is already set up)

set -e

export DATABASE_URL="postgres://postgres:postgres@localhost:5432/wot-statistics?sslmode=disable"
export TANKS_JSON_FILE="tanks_X.json"

echo "🚀 Quick seeding tanks data..."
go run cmd/seed/main.go
