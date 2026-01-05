package config

import "time"

type WargamingConfig struct {
	AppID   string
	BaseURL string
	Region  string
}

type DatabaseConfig struct {
	URL             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

type AppConfig struct {
	Wargaming *WargamingConfig
}

func GetWargamingConfig() *WargamingConfig {
	return &WargamingConfig{
		AppID:   "d5c27b088716f6a2ca4d043e6fe2ba91",
		BaseURL: "https://api.worldoftanks.asia/wot",
		Region:  "com",
	}
}

func GetAppConfig() *AppConfig {
	return &AppConfig{
		Wargaming: GetWargamingConfig(),
	}
}
