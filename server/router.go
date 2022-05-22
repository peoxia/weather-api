package server

import (
	"github.com/peoxia/weather-api/server/handler"
)

const v1API string = "/v1"

func (s *Server) setupRoutes() {
	s.Router.HandleFunc("/_healthz", handler.Healthz).Methods("GET").Name("Health")

	// API routes
	api := s.Router.PathPrefix(v1API).Subrouter()
	api.HandleFunc("/weather", handler.GetWeather(s.WeatherStack, s.OpenWeatherMap)).Methods("GET").Name("GetWeather")

}
