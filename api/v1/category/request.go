package category

import "github.com/zakiafada32/retail/modules/category"

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
