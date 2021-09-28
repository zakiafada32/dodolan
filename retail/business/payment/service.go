package payment

import (
	"errors"
	"log"

	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/utils"
)

type service struct {
	repository Repository
}

func NewPaymentService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) FindAll() ([]PaymentProvider, error) {
	payments, err := s.repository.FindAll()
	if err != nil {
		log.Println(err)
		return payments, errors.New(business.InternalServerError)
	}

	return payments, nil
}

func (s *service) CreateNew(provider PaymentProvider) error {
	err := utils.GetValidator().Struct(provider)
	if err != nil {
		log.Println(err)
		return err
	}
	err = s.repository.CreateNew(provider)
	if err != nil {
		log.Println(err)
		return errors.New(business.BadRequest)
	}
	return nil
}

func (s *service) Update(id uint32, name string, description string) (PaymentProvider, error) {
	payment, err := s.repository.Update(id, name, description)
	if err != nil {
		log.Println(err)
		return PaymentProvider{}, errors.New(business.BadRequest)
	}
	return payment, nil
}
