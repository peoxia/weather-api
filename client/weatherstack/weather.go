package weatherstack

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/peoxia/weather-api/client"
	"github.com/peoxia/weather-api/weather"
)

// GetWeather requests weather by city.
func (c *Client) GetWeather(city string) (*weather.Data, error) {

	reqData := client.HTTPRequestData{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s/current?access_key=%s&query=%s", c.Endpoint, c.AccessKey, city),
	}

	respBody, err := c.HTTPClient.RequestBytes(reqData)
	if err != nil {
		return nil, fmt.Errorf("error making request to get weather: %w", err)
	}

	var resp weatherResp
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling weather: %w", err)
	}

	data := weather.Data{
		WindSpeed:          resp.Current.WindSpeed,
		TemperatureDegrees: resp.Current.Temperature,
	}
	return &data, nil
}

type weatherResp struct {
	Current current `json:"current"`
}

type current struct {
	WindSpeed   int `json:"wind_speed"`
	Temperature int `json:"temperature"`
}
