package category

import "github.com/zakiafada32/retail/business/category"

type createNewCategoryRequestBody struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (req *createNewCategoryRequestBody) convertToCategoryBusiness() category.Category {
	return category.Category{
		Name:        req.Name,
		Description: req.Description,
	}
}

type updateCategoryRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
