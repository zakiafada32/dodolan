package courier

import "github.com/zakiafada32/retail/business/utils"

type service struct {
	repository Repository
}

func NewCourierService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateNewCourierProvider(courierProvider CourierProvider) error {
	err := utils.GetValidator().Struct(courierProvider)
	if err != nil {
		return err
	}
	return s.repository.CreateNewCourierProvider(courierProvider)
}
