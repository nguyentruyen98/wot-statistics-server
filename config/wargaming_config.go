package config

import (
	"os"
)

// WargamingConfig holds configuration for Wargaming API
// type WargamingConfig struct {
// 	AppID   string
// 	BaseURL string
// 	Region  string
// }

// GetWargamingConfig returns the Wargaming API configuration
// func GetWargamingConfig() *WargamingConfig {
// 	return &WargamingConfig{
// 		AppID:   getEnvOrDefault("WARGAMING_APP_ID", "d5c27b088716f6a2ca4d043e6fe2ba91"),
// 		BaseURL: getEnvOrDefault("WARGAMING_BASE_URL", "https://api.worldoftanks.asia/wot"),
// 		Region:  getEnvOrDefault("WARGAMING_REGION", "com"), // com, eu, ru, asia
// 	}
// }

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
