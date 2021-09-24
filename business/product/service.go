package product

import "github.com/zakiafada32/retail/business/utils"

type service struct {
	repository Repository
}

func NewProductService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateNewProduct(product Product) error {
	err := utils.GetValidator().Struct(product)
	if err != nil {
		return err
	}
	return s.repository.CreateNewProduct(product)
}
