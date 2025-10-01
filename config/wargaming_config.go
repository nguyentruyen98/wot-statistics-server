package config

import (
	"os"
)

// WargamingConfig holds configuration for Wargaming API
type WargamingConfig struct {
	AppID   string
	BaseURL string
	Region  string
}

// GetWargamingConfig returns the Wargaming API configuration
func GetWargamingConfig() *WargamingConfig {
	return &WargamingConfig{
		AppID:   getEnvOrDefault("WARGAMING_APP_ID", "your_app_id_here"),
		BaseURL: getEnvOrDefault("WARGAMING_BASE_URL", "https://api.worldoftanks.com/wot"),
		Region:  getEnvOrDefault("WARGAMING_REGION", "com"), // com, eu, ru, asia
	}
}

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
