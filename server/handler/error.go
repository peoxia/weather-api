package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// handleError logs the error and outputs the error message.
func handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Error(err.Error())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorJSON, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
	w.Write(errorJSON)
}
