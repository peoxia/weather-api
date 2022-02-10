package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peoxia/user-api/user"
	"github.com/peoxia/user-api/utils"
)

// DeleteUser deletes all data about a user by provided user ID.
// 		DELETE /api/v1/user/{user-id}
// 		Responds: 204, 400, 500
//		URI parameters:
//			id: The id of a user
func DeleteUser(storage user.Deleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := mux.Vars(r)["user-id"]
		if !utils.IsValidUUID(userID) {
			handleError(
				w,
				fmt.Errorf("error deleting user, ID %s is in the wrong format", userID),
				http.StatusBadRequest,
			)
			return
		}

		err := storage.DeleteUser(userID)
		if err != nil {
			handleError(
				w,
				fmt.Errorf("error deleting user with ID %s: %w", userID, err),
				http.StatusInternalServerError,
			)
			return
		}

		// Publish a message about a deleted user to PubSub.

		w.WriteHeader(http.StatusNoContent)
	}
}
