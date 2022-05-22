// Package config handles environment variables.
package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// Config contains environment variables.
type Config struct {
	Port                      string `envconfig:"PORT" default:"8000"`
	OpenWeatherMapAPIEndpoint string `envconfig:"OPEN_WEATHER_MAP_API_ENDPOINT" required:"true"`
	WeatherStackAPIEndpoint   string `envconfig:"WEATHER_STACK_API_ENDPOINT" required:"true"`
	WeatherStackAccessKey     string `envconfig:"WEATHER_STACK_ACCESS_KEY" required:"true"`
}

// LoadConfig reads environment variables and populates Config.
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found")
	}

	var c Config

	err := envconfig.Process("", &c)

	return &c, err
}
