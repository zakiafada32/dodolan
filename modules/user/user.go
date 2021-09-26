package user

import (
	"errors"
	"time"

	userBusiness "github.com/zakiafada32/retail/business/user"
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

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) CreateNewUser(user userBusiness.User) error {
	if err := repo.db.Where("email = ?", user.Email).First(&User{}).Error; err == nil {
		return errors.New("email already exist")
	}

	userData := convertToUserModel(user)

	if err := repo.db.Create(&userData).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) FindById(userId string) (userBusiness.User, error) {
	var userData User
	if err := repo.db.Where("id = ?", userId).First(&userData).Error; err != nil {
		return userBusiness.User{}, err
	}

	userBusiness := convertToUserBusiness(userData)
	return userBusiness, nil
}

func (repo *UserRepository) FindByEmail(email string) (userBusiness.User, error) {
	var user User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return userBusiness.User{}, err
	}

	userBusiness := convertToUserBusiness(user)
	return userBusiness, nil
}

func (repo *UserRepository) UpdateUser(userId, name, address string) (userBusiness.User, error) {
	var user User
	if err := repo.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return userBusiness.User{}, err
	}

	if err := repo.db.Model(&user).Updates(&User{Name: name, Address: address}).Error; err != nil {
		return userBusiness.User{}, err
	}

	userBusiness := convertToUserBusiness(user)
	return userBusiness, nil
}

func convertToUserModel(user userBusiness.User) User {
	return User{
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

func convertToUserBusiness(user User) userBusiness.User {
	return userBusiness.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		IsAdmin:   *user.IsAdmin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
