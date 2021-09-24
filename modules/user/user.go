package user

import (
	"errors"
	"time"

	"github.com/zakiafada32/retail/business/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type User struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	Address   string
	IsAdmin   *bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(user user.User) *User {
	return &User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		IsAdmin:   &user.IsAdmin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateNewUser(user user.User) error {
	if err := ur.db.Where("email = ?", user.Email).First(&User{}).Error; err == nil {
		return errors.New("email already exist")
	}

	userData := NewUser(user)

	if err := ur.db.Create(userData).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) FindByEmail(email string) (user.User, error) {
	var user user.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
