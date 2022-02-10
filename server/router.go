package server

import (
	"github.com/peoxia/user-api/server/handler"
)

const v1API string = "/api/v1"

func (s *Server) setupRoutes() {
	s.Router.HandleFunc("/_healthz", handler.Healthz).Methods("GET").Name("Health")

	// API routes
	api := s.Router.PathPrefix(v1API).Subrouter()
	api.HandleFunc("/users", handler.GetUsers(s.Database)).Methods("GET").Name("GetUsersList")
	api.HandleFunc("/users/{id}", handler.GetUser(s.Database)).Methods("GET").Name("GetUser")
	api.HandleFunc("/users/{id}", handler.CreateOrUpdateUser(s.Database)).Methods("POST").Name("CreateOrUpdateUser")
	api.HandleFunc("/users/{id}", handler.DeleteUser(s.Database)).Methods("DELETE").Name("DeleteUser")

}
