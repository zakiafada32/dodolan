package category

import (
	"errors"

	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/utils"
)

type service struct {
	repository Repository
}

func NewCategoryService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) CreateNew(category Category) error {
	err := utils.GetValidator().Struct(category)
	if err != nil {
		return errors.New(business.BadRequest)
	}
	err = s.repository.CreateNew(category)
	if err != nil {
		return errors.New(business.BadRequest)
	}
	return nil
}

func (s *service) FindAll() ([]Category, error) {
	categories, err := s.repository.FindAll()
	if err != nil {
		return []Category{}, errors.New(business.InternalServerError)
	}

	return categories, nil
}

func (s *service) FindById(id uint32) (Category, error) {
	productByCategory, err := s.repository.FindById(id)
	if err != nil {
		return Category{}, errors.New(business.NotFound)
	}

	return productByCategory, nil
}

func (s *service) Update(id uint32, name string, description string) (Category, error) {
	category, err := s.repository.Update(id, name, description)
	if err != nil {
		return Category{}, errors.New(business.BadRequest)
	}

	return category, nil
}
