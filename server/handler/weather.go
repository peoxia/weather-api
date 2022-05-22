package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/peoxia/weather-api/weather"
	log "github.com/sirupsen/logrus"
)

// GetWeather returns current weather at the requested location.
// Returns weather in Sydney if location is not specified.
// 		GET /v1/weather
// 		Responds: 200, 500
//		GET parameters:
//			city: The city
func GetWeather(primaryProvider weather.Fetcher, failoverProvider weather.Fetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()
		city := query.Get("city")
		if city == "" {
			city = "Sydney"
		}

		primaryChan := make(chan weatherResp, 1)
		failoverChan := make(chan weatherResp, 1)

		go func() {
			primaryChan <- getWeather(primaryProvider, city)
		}()

		go func() {
			failoverChan <- getWeather(failoverProvider, city)
		}()

		primaryResp := <-primaryChan
		failoverResp := <-failoverChan

		var collectedErrors error
		if primaryResp.err != nil {
			collectedErrors = fmt.Errorf("error fetching weather from primary provider: %w", primaryResp.err)
		}
		if failoverResp.err != nil {
			collectedErrors = fmt.Errorf("error fetching weather from failover provider: %v; previous error: %w",
				failoverResp.err,
				collectedErrors,
			)
		}

		var weather *weather.Data
		if primaryResp.err == nil {
			weather = primaryResp.data
		}
		if weather == nil && failoverResp.err == nil {
			weather = failoverResp.data
		}
		if weather == nil {
			handleError(
				w,
				fmt.Errorf("error fetching weather data: %w", collectedErrors),
				http.StatusInternalServerError,
			)
			return
		}
		if collectedErrors != nil {
			log.Error(collectedErrors)
		}

		response, err := json.Marshal(weather)
		if err != nil {
			handleError(
				w,
				fmt.Errorf("error marshalling weather data: %w", err),
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "public, max-age=3, stale-if-error=86400")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

type weatherResp struct {
	data *weather.Data
	err  error
}

func getWeather(fetcher weather.Fetcher, city string) weatherResp {
	data, err := fetcher.GetWeather(city)
	return weatherResp{
		data: data,
		err:  err,
	}
}
