package user_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/user"
	"github.com/zakiafada32/retail/business/user/mocks"
	"github.com/zakiafada32/retail/business/utils"
)

const (
	id       string = "id"
	name     string = "name"
	email    string = "name@email.com"
	password string = "password"
	address  string = "address"
	isAdmin  bool   = false
)

var (
	userService    user.Service
	userRepository mocks.Repository
	userData       user.User
	userDataRepo   user.User
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetCurrent(t *testing.T) {
	t.Run("Expect found the user with the same id and return the user data", func(t *testing.T) {
		userRepository.On("FindById", id).Return(userData, nil).Once()
		user, err := userService.GetCurrent(id)
		assert.Nil(t, err)
		assert.Equal(t, id, user.ID)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, address, user.Address)
		assert.Equal(t, password, user.Password)
	})

	t.Run("Expect internal server error when the user not found", func(t *testing.T) {
		userRepository.On("FindById", id).Return(user.User{}, errors.New("error")).Once()
		_, err := userService.GetCurrent(id)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(business.InternalServerError))
	})
}

func TestCreateNew(t *testing.T) {
	t.Run("Expect create new user and return nil", func(t *testing.T) {
		userRepository.On("CreateNew", mock.AnythingOfType("user.User")).Return(userData, nil).Once()
		user, err := userService.CreateNew(userData)
		assert.Nil(t, err)
		assert.Equal(t, user.Name, userData.Name)
		assert.Equal(t, user.Email, userData.Email)
		assert.Equal(t, user.Address, userData.Address)
	})

	t.Run("Expect bad request when the email already exist", func(t *testing.T) {
		userRepository.On("CreateNew", mock.AnythingOfType("user.User")).Return(user.User{}, errors.New("error")).Once()
		_, err := userService.CreateNew(userData)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(business.BadRequest))
	})
}

func TestLogin(t *testing.T) {
	t.Run("Expect user login success", func(t *testing.T) {
		userRepository.On("FindByEmail", email).Return(userDataRepo, nil).Once()
		token, err := userService.Login(userData.Email, userData.Password)
		assert.Nil(t, err)
		assert.IsType(t, "string", token)
	})

	t.Run("Expect user login faied, when email not found", func(t *testing.T) {
		userRepository.On("FindByEmail", email).Return(userDataRepo, errors.New("error")).Once()
		_, err := userService.Login(userData.Email, userData.Password)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(business.Unauthorized))
	})
}

func TestUpdate(t *testing.T) {

	t.Run("Expect update the user data success", func(t *testing.T) {
		userRepository.On("Update", id, name, address).Return(userDataRepo, nil).Once()
		user, err := userService.Update(id, name, address)
		assert.Nil(t, err)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, address, user.Address)
	})

	t.Run("Expect internal server error when the user not found", func(t *testing.T) {
		userRepository.On("Update", id, name, address).Return(user.User{}, errors.New("error")).Once()
		_, err := userService.Update(id, name, address)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(business.InternalServerError))
	})

}

func setup() {
	userData = user.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Address:  address,
		IsAdmin:  isAdmin,
	}

	hasingPassword, _ := utils.Hashing(password)

	userDataRepo = user.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: hasingPassword,
		Address:  address,
		IsAdmin:  isAdmin,
	}

	userService = user.NewUserService(&userRepository)
}
