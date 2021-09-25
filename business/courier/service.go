package courier

import "github.com/zakiafada32/retail/business/utils"

type service struct {
	repository Repository
}

func NewCourierService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) CreateNewCourierProvider(provider CourierProvider) error {
	err := utils.GetValidator().Struct(provider)
	if err != nil {
		return err
	}
	return s.repository.CreateNewCourierProvider(provider)
}
