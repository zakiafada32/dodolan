package modules

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string
	Name      string
	Email     string
	Password  string
	Address   string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
