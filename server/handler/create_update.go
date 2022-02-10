package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/peoxia/user-api/user"
	"github.com/peoxia/user-api/utils"
)

// CreateOrUpdateUser creates a user if it doesn't exist, and updates it
// if it does exist.
// 		POST /api/v1/users
// 		Responds: 200, 400, 500
//		Payload:
//			type Data struct {
//				ID        string     `json:"id"`
//				FirstName string     `json:"first_name,omitempty"`
//				LastName  string     `json:"last_name,omitempty"`
//				Nickname  string     `json:"nickname,omitempty"`
//				Password  string     `json:"password,omitempty"`
//				Email     string     `json:"email,omitempty"`
//				Country   string     `json:"country,omitempty"`
//				CreatedAt *time.Time `json:"created_at,omitempty"`
//				UpdatedAt *time.Time `json:"updated_at,omitempty"`
//			}

func CreateOrUpdateUser(storage user.CreatorUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Decode the request body
		var payload user.Data
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			handleError(
				w,
				fmt.Errorf("error decoding payload to create or update user: %w", err),
				http.StatusBadRequest,
			)
			return
		}

		// Validate payload
		if payload.ID != "" && !utils.IsValidUUID(payload.ID) {
			handleError(
				w,
				fmt.Errorf("error validating payload to create user, invalid ID"),
				http.StatusBadRequest,
			)
			return
		}

		userID, err := storage.CreateOrUpdateUser(payload)
		if err != nil {
			handleError(
				w,
				fmt.Errorf("error creating or updating user: %w", err),
				http.StatusInternalServerError,
			)
			return
		}

		// Publish a message about a created or updated user to PubSub here.

		response, err := json.Marshal(userID)
		if err != nil {
			handleError(
				w,
				fmt.Errorf("error marshalling create or update user response: %w", err),
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
