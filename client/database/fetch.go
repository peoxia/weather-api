package database

import (
	"fmt"
	"time"

	"github.com/peoxia/user-api/user"
)

// Declare prepareGetUserStmt() function here.

// GetUser is a mock function for fetching data about a user by provided ID.
func (c *Client) GetUser(id string) (*user.Data, error) {
	// Execute GetUser statement here.

	// Mocked responses
	if id == "d2a7924e-765f-4949-bc4c-219c956d0f8b" {
		createdAt := time.Date(2019, 10, 12, 7, 20, 50, 52, nil)
		return &user.Data{
			ID:        "d2a7924e-765f-4949-bc4c-219c956d0f8b",
			FirstName: "Alice",
			LastName:  "Bob",
			Nickname:  "AB123",
			Password:  "supersecurepassword",
			Email:     "alice@bob.com",
			Country:   "UK",
			CreatedAt: &createdAt,
			UpdatedAt: &createdAt,
		}, nil
	}
	return nil, fmt.Errorf("error could not find row with userUUID %s", id)
}

// Declare prepareGetUsersStmt() function here.

// GetUsers is a mock function for fetching a list of users by country,
// created date and page ID.
func (c *Client) GetUsers(pageID int, country string) ([]user.Data, error) {
	// Execute GetUsers statement here. It will select users by
	// requested parameters.

	// Mocked responses
	var users []user.Data
	if country == "UK" || country == "" {
		users = append(users, user.Data{
			ID:       "42624a20-0864-41a6-8be8-167461021454",
			Nickname: "John Lennon",
			Country:  "UK",
		})
	}
	if country == "UK" || country == "" {
		users = append(users, user.Data{
			ID:       "12f6a9c2-69fa-4736-9d02-1c771fce482f",
			Nickname: "Captain James Cook",
			Country:  "UK",
		})
	}
	if country == "Italy" || country == "" {
		users = append(users, user.Data{
			ID:       "201a924b-542b-4c2c-9221-140b2684c484",
			Nickname: "Galileo",
			Country:  "Italy",
		})
	}
	return users, nil
}
