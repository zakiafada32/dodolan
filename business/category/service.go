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

func (s *service) CreateNewCategory(category Category) error {
	err := utils.GetValidator().Struct(category)
	if err != nil {
		return errors.New(business.BadRequest)
	}
	return s.repository.CreateNewCategory(category)
}

func (s *service) FindAllCategory() ([]Category, error) {
	categories, err := s.repository.FindAllCategory()
	if err != nil {
		return []Category{}, errors.New(business.InternalServerError)
	}

	return categories, nil
}

func (s *service) FindCategoryById(categoryId uint32) (Category, error) {
	productByCategory, err := s.repository.FindCategoryById(categoryId)
	if err != nil {
		return Category{}, errors.New(business.BadRequest)
	}

	return productByCategory, nil
}

func (s *service) UpdateCategory(categoryId uint32, name string, description string) (Category, error) {
	category, err := s.repository.UpdateCategory(categoryId, name, description)
	if err != nil {
		return Category{}, errors.New(business.BadRequest)
	}

	return category, nil
}
