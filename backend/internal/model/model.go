package model

import (
	"regexp"
	"time"
)

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
	return u.ValidateUsername() && u.ValidateFirstname() && u.ValidateLastname() && u.ValidateEmail() && u.ValidateStatus() && u.ValidateDepartment()
}

func (u *User) ValidateUsername() bool {
	return len(u.Username) >= 3 && len(u.Username) <= 20 && regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(u.Username)
}

func (u *User) ValidateFirstname() bool {
	return len(u.Firstname) >= 3 && len(u.Firstname) <= 20 && regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(u.Firstname)
}

func (u *User) ValidateLastname() bool {
	return len(u.Lastname) >= 3 && len(u.Lastname) <= 20 && regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(u.Lastname)
}

func (u *User) ValidateEmail() bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(u.Email)
}

func (u *User) ValidateStatus() bool {
	return u.Status == "A" || u.Status == "I" || u.Status == "T"
}

func (u *User) ValidateDepartment() bool {
	return len(u.Department) >= 1 && len(u.Department) <= 20
}
