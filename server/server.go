// Package server provides functionality to easily set up an HTTTP server.
//
// The server holds all the clients it needs and they should be set up in the Create method.
//
// The HTTP routes and middleware are set up in the setupRouter method.
package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/peoxia/weather-api/client/openweathermap"
	"github.com/peoxia/weather-api/client/weatherstack"
	"github.com/peoxia/weather-api/config"

	log "github.com/sirupsen/logrus"
)

// Server holds the HTTP server, router, config and all clients.
type Server struct {
	Config         *config.Config
	OpenWeatherMap *openweathermap.Client
	WeatherStack   *weatherstack.Client
	HTTP           *http.Server
	Router         *mux.Router
}

// Create sets up the HTTP server, router and all clients.
// Returns an error if an error occurs.
func (s *Server) Create(ctx context.Context, config *config.Config) error {

	var OpenWeatherMapClient openweathermap.Client
	if err := OpenWeatherMapClient.Init(config); err != nil {
		return fmt.Errorf("error initializing Open Weather Map client: %w", err)
	}

	var WeatherStackClient weatherstack.Client
	if err := WeatherStackClient.Init(config); err != nil {
		return fmt.Errorf("error initializing Weather Stack client: %w", err)
	}

	s.Config = config
	s.OpenWeatherMap = &OpenWeatherMapClient
	s.WeatherStack = &WeatherStackClient
	s.Router = mux.NewRouter()
	s.HTTP = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Config.Port),
		Handler: s.Router,
	}

	s.setupRoutes()

	return nil
}

// Serve tells the server to start listening and serve HTTP requests.
// It also makes sure that the server gracefully shuts down on exit.
// Returns an error if an error occurs.
func (s *Server) Serve(ctx context.Context) error {
	idleConnsClosed := make(chan struct{}) // this is used to signal that we can not exit
	go func(ctx context.Context, s *Server) {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		<-stop

		log.Info("Shutdown signal received")

		if err := s.HTTP.Shutdown(ctx); err != nil {
			log.Error(err.Error())
		}

		close(idleConnsClosed) // call close to say we can now exit the function
	}(ctx, s)

	log.Infof("Ready at: %s", s.Config.Port)

	if err := s.HTTP.ListenAndServe(); err != http.ErrServerClosed {
		log.Error(err.Error())
	}
	<-idleConnsClosed // this will block until close is called

	return nil
}
