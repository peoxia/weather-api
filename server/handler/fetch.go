package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/peoxia/user-api/user"
	"github.com/peoxia/user-api/utils"
)

// GetUser returns data about a user by provided ID.
// 		GET /api/v1/users/{user-id}
// 		Responds: 200, 400, 500
//		URI parameters:
//			id: The id of a user
func GetUser(storage user.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := mux.Vars(r)["user-id"]
		if !utils.IsValidUUID(userID) {
			handleError(
				w,
				fmt.Errorf("error getting user, ID %s is in the wrong format", userID),
				http.StatusBadRequest,
			)
		}

		user, err := storage.GetUser(userID)
		if err != nil {
			handleError(
				w,
				fmt.Errorf("error getting user with ID %s: %w", userID, err),
				http.StatusInternalServerError,
			)
		}

		response, err := json.Marshal(user)
		if err != nil {
			handleError(
				w,
				fmt.Errorf("error marshalling user data: %w", err),
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

// GetUserList returns a requested page of users list by country and
// created date if provided.
// 		GET /api/v1/users
// 		Responds: 200, 400, 500
//		GET parameters:
//			pageId: The id of a page
//			country: The country of users to lookup
func GetUsers(storage user.ListGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get parameters
		query := r.URL.Query()
		pageIDRaw := query.Get("pageId")
		pageID, err := strconv.Atoi(pageIDRaw)
		if err != nil {
			handleError(
				w,
				fmt.Errorf("error getting users list, pageID has wrong format: %w", err),
				http.StatusBadRequest,
			)
		}
		country := query.Get("country")

		users, err := storage.GetUsers(pageID, country)
		if err != nil {
			handleError(
				w,
				fmt.Errorf("error getting users list: %w", err),
				http.StatusInternalServerError,
			)
		}

		response, err := json.Marshal(users)
		if err != nil {
			handleError(
				w,
				fmt.Errorf("error marshalling users list: %w", err),
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
