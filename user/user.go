package user

import (
	"time"
)

type Data struct {
	ID        string     `json:"id"`
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	Nickname  string     `json:"nickname,omitempty"`
	Password  string     `json:"password,omitempty"`
	Email     string     `json:"email,omitempty"`
	Country   string     `json:"country,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
