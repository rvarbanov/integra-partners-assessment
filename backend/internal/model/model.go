package model

import "time"

// TODO: normalize data - department
// create department table and map department name to id
type User struct {
	ID         int       `db:"user_id" json:"user_id"`
	Username   string    `db:"user_name" json:"user_name,omitempty"`
	Firstname  string    `db:"first_name" json:"first_name,omitempty"`
	Lastname   string    `db:"last_name" json:"last_name,omitempty"`
	Email      string    `db:"email" json:"email,omitempty"`
	Status     string    `db:"user_status" json:"user_status,omitempty"`
	Department string    `db:"department" json:"department,omitempty"`
	CreatedAt  time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

type Response struct {
	Data  any `json:"data,omitempty"`
	Error any `json:"error,omitempty"`
}

func (u *User) IsValid() bool {
	// TODO: validation for the data
	// - username - must be between 3 and 20 characters, alphanumeric
	// - firstname - must be between 1 and 20 characters, alphabetic
	// - lastname - must be between 1 and 20 characters, alphabetic
	// - email - must be a valid email address
	// - status - must be "A", "I", or "T"
	// - department - must be one from a list

	return true
}
