package config

type WargamingConfig struct {
	AppID   string
	BaseURL string
	Region  string
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
