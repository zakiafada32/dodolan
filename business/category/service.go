package category

import "github.com/zakiafada32/retail/business/utils"

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
		return err
	}
	return s.repository.CreateNewCategory(category)
}
