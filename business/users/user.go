package users

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Address   string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(id string, name string, email string, password string, isAdmin bool) User {
	return User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		IsAdmin:   isAdmin,
		CreatedAt: time.Now(),
	}
}
