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

func (s *service) FindById(userId string, orderId uint32) (Order, error) {
	order, err := s.repository.FindById(userId, orderId)
	if err != nil {
		log.Println(err)
		return Order{}, errors.New(business.NotFound)
	}

	return order, nil
}

func (s *service) FindAll(userId string) ([]Order, error) {
	orders, err := s.repository.FindAll(userId)
	if err != nil {
		log.Println(err)
		return []Order{}, errors.New(business.NotFound)
	}

	return orders, nil
}

func (s *service) Payment(userId string, orderId uint32, totalAmount uint64) error {
	err := s.repository.Payment(userId, orderId, totalAmount)
	if err != nil {
		log.Println(err)
		return errors.New(business.BadRequest)
	}

	return nil
}

func (s *service) Courier(userId string, orderId uint32) error {
	err := s.repository.Courier(userId, orderId)
	if err != nil {
		log.Println(err)
		return errors.New(business.BadRequest)
	}

	return nil
}
