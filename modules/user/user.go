package user

import (
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
	IsAdmin   bool
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
		user.IsAdmin,
		user.CreatedAt,
		user.UpdatedAt,
	}
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) CreateNewUser(user user.User) error {
	if err := repo.db.Where("email = ?", user.Email).First(&User{}).Error; err != nil {
		log.Println("wokwokwokw")
		return err
	}

	userData := NewUser(user)
	userData.ID = utils.GenerateID()

	if err := repo.db.Create(userData).Error; err != nil {
		return err
	}

	return nil
}
