package order

import (
	"errors"

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
		return Order{}, errors.New(business.NotFound)
	}

	return order, nil
}

func (s *service) FindAll() ([]Order, error) {
	orders, err := s.repository.FindAll()
	if err != nil {
		return []Order{}, errors.New(business.NotFound)
	}

	return orders, nil
}
