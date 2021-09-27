package order

import (
	"errors"
	"log"

	"github.com/zakiafada32/retail/business"
)

type service struct {
	repository Repository
}

func NewOrderService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) FindById(orderId uint32) (Order, error) {
	order, err := s.repository.FindById(orderId)
	if err != nil {
		log.Println(err)
		return Order{}, errors.New(business.NotFound)
	}

	return order, nil
}

func (s *service) FindAll() ([]Order, error) {
	orders, err := s.repository.FindAll()
	if err != nil {
		log.Println(err)
		return []Order{}, errors.New(business.NotFound)
	}

	return orders, nil
}

func (s *service) Payment(orderId uint32, totalAmount uint64) error {
	err := s.repository.Payment(orderId, totalAmount)
	if err != nil {
		log.Println(err)
		return errors.New(business.BadRequest)
	}

	return nil
}
