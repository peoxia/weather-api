// Package openweathermap contains a client and methods for communicating
// with Open Weather Map.
package openweathermap

import (
	"time"

	"github.com/peoxia/weather-api/client"
	"github.com/peoxia/weather-api/config"
)

// Client holds the HTTP client and endpoint information.
type Client struct {
	Endpoint   string
	HTTPClient client.HTTPClient
}

// Init sets up a new Open Weather Map client.
func (c *Client) Init(config *config.Config) error {
	timeout := 3 * time.Second
	c.Endpoint = config.OpenWeatherMapAPIEndpoint
	c.HTTPClient = client.NewHTTPClient(client.Parameters{Timeout: &timeout})
	return nil
}
