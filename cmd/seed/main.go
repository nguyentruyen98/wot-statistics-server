package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"wot-statistics-server/config"
	"wot-statistics-server/domain"
	"wot-statistics-server/internal/infrastructure/postgres"
	"wot-statistics-server/repository"
	"wot-statistics-server/usecase"
)

func main() {
	fmt.Println("🚀 Starting tank data seeder...")

	// Get database configuration
	dbConfig := &config.DatabaseConfig{
		URL:             getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/wot-statistics?sslmode=disable"),
		MaxOpenConns:    25,
		MaxIdleConns:    25,
		ConnMaxLifetime: 5 * time.Minute,
		ConnMaxIdleTime: 10 * time.Minute,
	}

	// Connect to database
	fmt.Println("📡 Connecting to database...")
	db, err := postgres.NewDB(dbConfig)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer db.Close()

	fmt.Println("✅ Database connected successfully")

	// Get app config
	appConfig := config.GetAppConfig()

	// Initialize tank repository and usecase
	tankRepo := repository.NewTankRepository(db, appConfig)
	profileRepo := repository.NewProfileRepository(db)
	tankUseCase := usecase.NewTankUseCase(tankRepo, profileRepo)

	// Initialize module repository and usecase
	moduleRepo := repository.NewModuleRepository(db)
	moduleUseCase := usecase.NewModuleUseCase(moduleRepo, profileRepo)

	// Import tanks
	tanksFile := getEnv("TANKS_JSON_FILE", "tanks_X.json")
	if err := importTanks(tankUseCase, tanksFile); err != nil {
		log.Fatalf("❌ Failed to import tanks: %v", err)
	}

	// Import modules
	modulesFile := getEnv("MODULES_JSON_FILE", "module.json")
	if err := importModules(moduleUseCase, modulesFile); err != nil {
		log.Fatalf("❌ Failed to import modules: %v", err)
	}

	// Show statistics
	fmt.Println("\n📊 Database Statistics:")
	showTankStatistics(tankUseCase)
	showModuleStatistics(moduleUseCase)
}

// importTanks imports tank data from JSON file
func importTanks(uc domain.TankUseCase, jsonFilePath string) error {
	// Check if file exists
	if _, err := os.Stat(jsonFilePath); os.IsNotExist(err) {
		return fmt.Errorf("JSON file not found: %s", jsonFilePath)
	}

	// Get absolute path
	absPath, err := filepath.Abs(jsonFilePath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	fmt.Printf("\n📄 Reading tanks data from: %s\n", absPath)

	// Import tanks from JSON
	startTime := time.Now()
	count, err := uc.ImportTanksFromJSON(absPath)
	elapsed := time.Since(startTime)

	if err != nil {
		return err
	}

	// Print success message
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("✅ Successfully imported %d tanks!\n", count)
	fmt.Printf("⏱️  Time taken: %v\n", elapsed)
	fmt.Println(strings.Repeat("=", 60))

	return nil
}

// importModules imports module data from JSON file
func importModules(uc domain.ModuleUseCase, jsonFilePath string) error {
	// Check if file exists
	if _, err := os.Stat(jsonFilePath); os.IsNotExist(err) {
		return fmt.Errorf("JSON file not found: %s", jsonFilePath)
	}

	// Get absolute path
	absPath, err := filepath.Abs(jsonFilePath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	fmt.Printf("\n📄 Reading modules data from: %s\n", absPath)

	// Import modules from JSON
	startTime := time.Now()
	count, err := uc.ImportModulesFromJSON(absPath)
	elapsed := time.Since(startTime)

	if err != nil {
		return err
	}

	// Print success message
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("✅ Successfully imported %d modules!\n", count)
	fmt.Printf("⏱️  Time taken: %v\n", elapsed)
	fmt.Println(strings.Repeat("=", 60))

	return nil
}

// showTankStatistics displays statistics about imported tanks
func showTankStatistics(uc domain.TankUseCase) {
	// Get all tanks
	tanks, err := uc.GetTanks()
	if err != nil {
		fmt.Printf("⚠️  Failed to get tank statistics: %v\n", err)
		return
	}

	// Count by nation
	nationCount := make(map[string]int)
	tierCount := make(map[int]int)
	typeCount := make(map[string]int)
	premiumCount := 0

	for _, tank := range tanks {
		nationCount[tank.Nation]++
		tierCount[tank.Tier]++
		typeCount[tank.Type]++
		if tank.IsPremium {
			premiumCount++
		}
	}

	fmt.Println("\n🎮 TANKS:")
	fmt.Printf("   Total tanks: %d\n", len(tanks))
	fmt.Printf("   Premium tanks: %d\n", premiumCount)

	fmt.Println("\n   By Nation:")
	for nation, count := range nationCount {
		fmt.Printf("      %s: %d\n", nation, count)
	}

	fmt.Println("\n   By Tier:")
	for tier := 1; tier <= 10; tier++ {
		if count, ok := tierCount[tier]; ok {
			fmt.Printf("      Tier %d: %d\n", tier, count)
		}
	}

	fmt.Println("\n   By Type:")
	for tankType, count := range typeCount {
		fmt.Printf("      %s: %d\n", tankType, count)
	}
}

// showModuleStatistics displays statistics about imported modules
func showModuleStatistics(uc domain.ModuleUseCase) {
	// Get all modules
	modules, err := uc.GetModules()
	if err != nil {
		fmt.Printf("⚠️  Failed to get module statistics: %v\n", err)
		return
	}

	// Count by type
	typeCount := make(map[string]int)
	nationCount := make(map[string]int)
	tierCount := make(map[int]int)

	for _, module := range modules {
		typeCount[module.Type]++
		nationCount[module.Nation]++
		tierCount[module.Tier]++
	}

	fmt.Println("\n🔧 MODULES:")
	fmt.Printf("   Total modules: %d\n", len(modules))

	fmt.Println("\n   By Type:")
	for moduleType, count := range typeCount {
		fmt.Printf("      %s: %d\n", moduleType, count)
	}

	fmt.Println("\n   By Nation:")
	for nation, count := range nationCount {
		fmt.Printf("      %s: %d\n", nation, count)
	}

	fmt.Println("\n   By Tier:")
	for tier := 1; tier <= 10; tier++ {
		if count, ok := tierCount[tier]; ok {
			fmt.Printf("      Tier %d: %d\n", tier, count)
		}
	}
}

// getEnv gets environment variable with a default fallback
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
