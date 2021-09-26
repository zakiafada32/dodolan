package product

import (
	"errors"

	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/utils"
)

type service struct {
	repository Repository
}

func NewProductService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) FindById(id uint32) (ProductAtt, error) {
	product, err := s.repository.FindById(id)
	if err != nil {
		return ProductAtt{}, errors.New(business.NotFound)
	}
	return product, nil
}

func (s *service) FindAll() ([]ProductAtt, error) {
	products, err := s.repository.FindAll()
	if err != nil {
		return []ProductAtt{}, errors.New(business.BadRequest)
	}
	return products, nil
}

func (s *service) CreateNew(product Product) error {
	err := utils.GetValidator().Struct(product)
	if err != nil {
		return err
	}
	err = s.repository.CreateNew(product)
	if err != nil {
		return errors.New(business.BadRequest)
	}
	return nil
}

func (s *service) Update(id uint32, updateData ProductUpdate) (ProductAtt, error) {
	product, err := s.repository.Update(id, updateData)
	if err != nil {
		return ProductAtt{}, errors.New(business.BadRequest)
	}

	return product, nil
}
