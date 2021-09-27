package category

import "github.com/zakiafada32/retail/business/category"

type categoryRequestBody struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (req *categoryRequestBody) convertToCategoryBusiness() category.Category {
	return category.Category{
		Name:        req.Name,
		Description: req.Description,
	}
}
