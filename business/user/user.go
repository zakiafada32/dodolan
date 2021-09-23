package user

import "time"

type User struct {
	ID        string
	Name      string `validate:"required"`
	Email     string `validate:"required"`
	Password  string `validate:"required"`
	Address   string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
