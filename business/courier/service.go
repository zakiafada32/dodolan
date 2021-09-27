package courier

import (
	"errors"
	"log"

	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/utils"
)

type service struct {
	repository Repository
}

func NewCourierService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) FindAll() ([]CourierProvider, error) {
	couriers, err := s.repository.FindAll()
	if err != nil {
		log.Println(err)
		return []CourierProvider{}, errors.New(business.InternalServerError)
	}

	return couriers, nil
}

func (s *service) CreateNew(provider CourierProvider) error {
	err := utils.GetValidator().Struct(provider)
	if err != nil {
		log.Println(err)
		return errors.New(business.BadRequest)
	}
	return s.repository.CreateNew(provider)
}

func (s *service) Update(id uint32, name string, description string) (CourierProvider, error) {
	courier, err := s.repository.Update(id, name, description)
	if err != nil {
		log.Println(err)
		return CourierProvider{}, errors.New(business.BadRequest)
	}
	return courier, nil
}
