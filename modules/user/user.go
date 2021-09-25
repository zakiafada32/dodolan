package user

import (
	"errors"
	"time"

	"github.com/zakiafada32/retail/business/user"
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

func (ur *UserRepository) CreateNewUser(user userBusiness.User) error {
	if err := ur.db.Where("email = ?", user.Email).First(&User{}).Error; err == nil {
		return errors.New("email already exist")
	}

	userData := convertToUserModel(user)

	if err := ur.db.Create(&userData).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) FindById(userId string) (userBusiness.User, error) {
	var user User
	if err := ur.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return userBusiness.User{}, err
	}

	userBusiness := convertToUserBusiness(user)
	return userBusiness, nil
}

func (ur *UserRepository) FindByEmail(email string) (userBusiness.User, error) {
	var user User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return userBusiness.User{}, err
	}

	userBusiness := convertToUserBusiness(user)
	return userBusiness, nil
}

func (ur *UserRepository) UpdateUser(userId string, updateData user.UpdateUser) (userBusiness.User, error) {
	var user User
	if err := ur.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return userBusiness.User{}, err
	}

	if err := ur.db.Model(&user).Updates(&User{Name: updateData.Name, Address: updateData.Address}).Error; err != nil {
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
