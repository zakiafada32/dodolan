package user

import "github.com/zakiafada32/retail/business/utils"

type service struct {
	repository Repository
}

func NewUserService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateNewUser(user User) error {
	hashingPassword, err := utils.Hashing(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashingPassword
	return s.repository.CreateNewUser(user)
}

func (s *service) Login(email string, password string) (string, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = utils.CompareHash(user.Password, password)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID, user.IsAdmin)
	return token, err
}
