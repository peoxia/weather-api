// Package weather contains data structures, interfaces and unit
// conversion tools.
package weather

// Data contains data about weather at a specific location.
type Data struct {
	WindSpeed          int `json:"wind_speed"`
	TemperatureDegrees int `json:"temperature_degrees"`
}

type Fetcher interface {
	GetWeather(city string) (*Data, error)
}
