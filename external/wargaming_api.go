package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"wot-statistics-server/config"
)

// WargamingAPI represents the Wargaming API client
type WargamingAPI struct {
	BaseURL    string
	AppID      string
	HTTPClient *http.Client
}

// APIResponse represents a generic API response structure
type APIResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  *APIError   `json:"error,omitempty"`
}

// APIError represents an error from the Wargaming API
type APIError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Value   string `json:"value"`
}

// NewWargamingAPI creates a new Wargaming API client
func NewWargamingAPI(appConfig *config.AppConfig) *WargamingAPI {
	return &WargamingAPI{
		BaseURL: appConfig.Wargaming.BaseURL,
		AppID:   appConfig.Wargaming.AppID,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// "https://api.worldoftanks.asia/wot/encyclopedia/vehicles/?application_id=d5c27b088716f6a2ca4d043e6fe2ba91&language=vi&fields=description&tier=1"
// "https://api.worldoftanks.asia/wot/encyclopedia/vehicles/?application_id=d5c27b088716f6a2ca4d043e6fe2ba91&fields=description&language=vi&tier=1"
// "https://api.worldoftanks.asia/wot/encyclopedia/vehicles/?application_id=d5c27b088716f6a2ca4d043e6fe2ba91&limit=100&fields=tank_id,name,type,tier"
// makeRequest performs an HTTP GET request to the Wargaming API
func (w *WargamingAPI) makeRequest(endpoint string, params map[string]string) (*APIResponse, error) {
	// Build URL with parameters
	url := fmt.Sprintf("%s%s?application_id=%s", w.BaseURL, endpoint, w.AppID)

	// Add additional parameters
	for key, value := range params {
		url += fmt.Sprintf("&%s=%s", key, value)
	}
	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers
	req.Header.Set("User-Agent", "WoT-Statistics-Server/1.0")
	req.Header.Set("Accept", "application/json")

	// Make HTTP request
	resp, err := w.HTTPClient.Do(req)
	// Make HTTP request
	// resp, err := w.HTTPClient.Get(url)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to make request: %w", err)
	// }
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON response
	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Check for API errors
	if apiResp.Status != "ok" {
		if apiResp.Error != nil {
			return nil, fmt.Errorf("API error: %s (code: %d)", apiResp.Error.Message, apiResp.Error.Code)
		}
		return nil, fmt.Errorf("API returned status: %s", apiResp.Status)
	}

	return &apiResp, nil
}

// GetTanksList retrieves the list of tanks from Wargaming API
func (w *WargamingAPI) GetTanksList(params map[string]string) (*APIResponse, error) {
	return w.makeRequest("/encyclopedia/vehicles/", params)
}

// GetPlayerInfo retrieves player information from Wargaming API
func (w *WargamingAPI) GetPlayerInfo(accountID string) (*APIResponse, error) {
	params := map[string]string{
		"account_id": accountID,
	}
	return w.makeRequest("/account/info/", params)
}

// SearchPlayers searches for players by nickname
func (w *WargamingAPI) SearchPlayers(search string) (*APIResponse, error) {
	params := map[string]string{
		"search": search,
	}
	return w.makeRequest("/account/list/", params)
}
