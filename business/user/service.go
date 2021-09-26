package user

import (
	"errors"

	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/utils"
)

type service struct {
	repository Repository
}

func NewUserService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetCurrentUser(userId string) (User, error) {
	user, err := s.repository.FindById(userId)
	if err != nil {
		return User{}, errors.New(business.InternalServerError)
	}

	return user, nil
}

func (s *service) CreateNewUser(user User) error {
	err := utils.GetValidator().Struct(user)
	if err != nil {
		return errors.New(business.BadRequest)
	}
	user.ID = utils.GenerateID()
	hashingPassword, err := utils.Hashing(user.Password)
	if err != nil {
		return errors.New(business.InternalServerError)
	}

	user.Password = hashingPassword
	return s.repository.CreateNewUser(user)
}

func (s *service) Login(email string, password string) (string, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return "", errors.New(business.Unauthorized)
	}

	err = utils.CompareHash(user.Password, password)
	if err != nil {
		return "", errors.New(business.Unauthorized)
	}

	token, err := utils.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		return "", errors.New(business.InternalServerError)
	}
	return token, nil
}

func (s *service) UpdateUser(userId, name, address string) (User, error) {
	user, err := s.repository.UpdateUser(userId, name, address)
	if err != nil {
		return User{}, errors.New(business.InternalServerError)
	}

	return user, nil
}
