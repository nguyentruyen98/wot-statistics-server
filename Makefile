# Variables
MAIN_PATH=./cmd/api


# Color output
CYAN=\033[0;36m
NC=\033[0m # No Color

## run: Run the application locally
run:
	@echo "$(CYAN)Running application...$(NC)"
	@go run $(MAIN_PATH)/main.go