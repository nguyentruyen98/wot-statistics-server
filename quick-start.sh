#!/bin/bash

# Quick Start Script - Import Tank Data
# This script helps you quickly import tank data into the database

set -e

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
RED='\033[0;31m'
NC='\033[0m'

clear

echo -e "${CYAN}"
cat << "EOF"
╔══════════════════════════════════════════════════════════╗
║                                                          ║
║       WoT Statistics - Tank Data Import Tool           ║
║                                                          ║
╚══════════════════════════════════════════════════════════╝
EOF
echo -e "${NC}"

# Menu
echo -e "${GREEN}Chọn phương thức import:${NC}"
echo ""
echo "  1) 🚀 Auto Setup + Import (Khuyến nghị lần đầu)"
echo "  2) ⚡ Quick Import (Database đã sẵn sàng)"
echo "  3) 🔧 Manual Step by Step"
echo "  4) 🔄 Reset Database + Import lại"
echo "  5) ❌ Thoát"
echo ""
read -p "Nhập lựa chọn (1-5): " choice

case $choice in
    1)
        echo -e "${YELLOW}Chạy auto setup + import...${NC}"
        make seed
        ;;
    2)
        echo -e "${YELLOW}Chạy quick import...${NC}"
        make seed-quick
        ;;
    3)
        echo -e "${CYAN}=== Manual Step by Step ===${NC}"
        echo ""
        
        echo -e "${YELLOW}Bước 1: Khởi động PostgreSQL...${NC}"
        docker-compose up -d postgres
        echo -e "${GREEN}✓ PostgreSQL started${NC}"
        sleep 3
        
        echo ""
        echo -e "${YELLOW}Bước 2: Tạo database...${NC}"
        docker exec wot-statistics-postgres psql -U postgres -c "CREATE DATABASE \"wot-statistics\";" 2>/dev/null || echo "Database đã tồn tại"
        echo -e "${GREEN}✓ Database ready${NC}"
        
        echo ""
        echo -e "${YELLOW}Bước 3: Chạy migrations...${NC}"
        docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/001_create_tanks_table.up.sql 2>&1 | grep -v "already exists" || true
        echo -e "${GREEN}✓ Migrations completed${NC}"
        
        echo ""
        echo -e "${YELLOW}Bước 4: Import data...${NC}"
        DATABASE_URL="postgres://postgres:postgres@localhost:5432/wot-statistics?sslmode=disable" \
        TANKS_JSON_FILE="tanks_X.json" \
        go run cmd/seed/main.go
        ;;
    4)
        echo -e "${YELLOW}Reset database và import lại...${NC}"
        make db-reset
        echo ""
        make seed-quick
        ;;
    5)
        echo -e "${CYAN}Tạm biệt!${NC}"
        exit 0
        ;;
    *)
        echo -e "${RED}Lựa chọn không hợp lệ!${NC}"
        exit 1
        ;;
esac

echo ""
echo -e "${GREEN}╔══════════════════════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║              Import hoàn thành!                         ║${NC}"
echo -e "${GREEN}╚══════════════════════════════════════════════════════════╝${NC}"
echo ""
echo -e "${CYAN}Kiểm tra data:${NC}"
echo "  • Open psql:    make db-psql"
echo "  • Start API:    make run"
echo "  • Test API:     curl http://localhost:8080/api/tanks"
echo ""
