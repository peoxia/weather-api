// Package database is a mock client of a MySQL database that would
// store user data.
package database

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // load the driver
	"github.com/jmoiron/sqlx"

	"github.com/peoxia/user-api/config"
)

// Client holds the database client and prepared statements.
type Client struct {
	Timeout                time.Duration
	DB                     *sqlx.DB
	GetUserStmt            *sqlx.Stmt
	GetUsersStmt           *sqlx.Stmt
	CreateOrUpdateUserStmt *sqlx.Stmt
	DeleteUserStmt         *sqlx.Stmt
}

// InitMock initializes mocked database client.
func (c *Client) InitMock(config *config.Config) error {
	// Connect to the database and set connection limits from config here.

	// Prepare statements here.

	return nil
}

// Close closes the database connection and statements.
func (c *Client) Close() error {
	// Close statements here.

	// Close connection with database here.

	return nil
}
