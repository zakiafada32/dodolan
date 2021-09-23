package user

import (
	"errors"
	"log"
	"time"

	"github.com/zakiafada32/retail/business/user"
	"github.com/zakiafada32/retail/modules/utils"
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
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.Address,
		&user.IsAdmin,
		user.CreatedAt,
		user.UpdatedAt,
	}
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateNewUser(user user.User) error {
	if err := u.db.Where("email = ?", user.Email).First(&User{}).Error; err == nil {
		return errors.New("email already exist")
	}

	userData := NewUser(user)
	log.Println(userData)
	userData.ID = utils.GenerateID()

	if err := u.db.Create(userData).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) FindByEmail(email string) (user.User, error) {
	var user user.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
