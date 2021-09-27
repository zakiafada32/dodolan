package cart

import (
	"errors"

	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/utils"
)

type service struct {
	repository Repository
}

func NewCartService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Update(userId string, cartItem CartItem) error {
	err := utils.GetValidator().Struct(cartItem)
	if err != nil {
		return errors.New(business.BadRequest)
	}

	err = s.repository.Update(userId, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		return errors.New(business.BadRequest)
	}

	return nil
}

func (s *service) FindAll(userId string) (Cart, error) {
	cartItem, err := s.repository.FindAll(userId)
	if err != nil {
		return Cart{}, errors.New(business.InternalServerError)
	}
	cart := Cart{}
	cart.Items = cartItem
	for _, item := range cartItem {
		cart.TotalAmount += uint64(item.TotalAmount)
	}
	return cart, nil
}

func (s *service) DeleteCartItem(userId string, productsId []uint32) error {
	err := s.repository.DeleteCartItem(userId, productsId)
	if err != nil {
		return errors.New(business.BadRequest)
	}

	return nil
}

func (s *service) Checkout(userId string, paymentId uint32, courierId uint32) error {
	cartItem, err := s.repository.FindAll(userId)
	if err != nil {
		return errors.New(business.InternalServerError)
	}
	if len(cartItem) == 0 {
		return errors.New(business.BadRequest)
	}
	productsId := make([]uint32, len(cartItem))
	cart := Cart{}
	cart.Items = cartItem
	for i, item := range cartItem {
		cart.TotalAmount += uint64(item.TotalAmount)
		productsId[i] = item.Product.ID
	}

	err = s.repository.Checkout(userId, paymentId, courierId, cart)
	if err != nil {
		return errors.New(business.InternalServerError)
	}

	err = s.repository.DeleteCartItem(userId, productsId)
	if err != nil {
		return errors.New(business.BadRequest)
	}

	return nil
}
