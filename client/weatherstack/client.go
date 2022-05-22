// Package weatherstack contains a client and methods for communicating
// with Weather Stack.
package weatherstack

import (
	"time"

	"github.com/peoxia/weather-api/client"
	"github.com/peoxia/weather-api/config"
)

// Client holds the HTTP client and endpoint information.
type Client struct {
	Endpoint   string
	AccessKey  string
	HTTPClient client.HTTPClient
}

// Init sets up a new Weather Stack client.
func (c *Client) Init(config *config.Config) error {
	timeout := 3 * time.Second
	c.Endpoint = config.WeatherStackAPIEndpoint
	c.AccessKey = config.WeatherStackAccessKey
	c.HTTPClient = client.NewHTTPClient(client.Parameters{Timeout: &timeout})
	return nil
}
