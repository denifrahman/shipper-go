package shipper

// Configuration is struct for API config.
type Configuration struct {
	APIKey  string
	BaseURL string
}

// Conf is instance of Configuration struct.
var Conf = Configuration{
	BaseURL: "https://merchant-api-sandbox.shipper.id/v3",
}

// SetAPIKey sets the API Key for API Call.
func (c *Configuration) SetAPIKey(token string) *Configuration {
	c.APIKey = token

	return c
}

// SetProductionMode sets the API base URL to production.
func (c *Configuration) SetProductionMode(productionMode bool) {
	if productionMode {
		c.BaseURL = "https://merchant-api.shipper.id/v3"
	}
}
