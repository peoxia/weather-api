package openweathermap

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
		URL:    fmt.Sprintf("%s/data/2.5/weather?q=%s,AU&appid=2326504fb9b100bee21400190e4dbe6d", c.Endpoint, city),
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
		WindSpeed:          int(resp.Wind.Speed),
		TemperatureDegrees: int(weather.KelvinToCelcius(resp.Main.Temp)),
	}
	return &data, nil
}

type weatherResp struct {
	Main main `json:"main"`
	Wind wind `json:"wind"`
}

type main struct {
	Temp float64 `json:"temp"`
}

type wind struct {
	Speed float64 `json:"speed"`
}
