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

func (s *service) UpdateCartItem(userId string, cartItem CartItem) error {
	err := utils.GetValidator().Struct(cartItem)
	if err != nil {
		return errors.New(business.BadRequest)
	}

	err = s.repository.UpdateCartItem(userId, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		return errors.New(business.BadRequest)
	}

	return nil
}
