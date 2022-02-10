package database

import (
	"fmt"

	"github.com/peoxia/user-api/user"
	"github.com/peoxia/user-api/utils"
)

// Declare prepareCreateOrUpdateUserStmt() function here.

// CreateOrUpdateUser is a mock function for creating a new user or
// updating user data. Returns ID of the user that was created or
// updated.
func (c *Client) CreateOrUpdateUser(user user.Data) (string, error) {
	// Execute CreateOrUpdateUser statement here.

	if user.ID == "" {
		userID, err := utils.GenerateUUID()
		if err != nil {
			return "", fmt.Errorf("error generating ID for a new user")
		}
		user.ID = userID
	}

	return user.ID, nil
}
