package category

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/business/category"
)

type CategoryController struct {
	service category.Service
}

func NewCategoryController(service category.Service) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

func (pc *CategoryController) CreateNewCategory(c echo.Context) error {
	var body createNewCategoryRequestBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	if err := c.Validate(&body); err != nil {
		return err
	}

	if err := pc.service.CreateNewCategory(category.Category(body.convertToCategoryBusiness())); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"code":    "new_category_created",
		"message": "new category has been created successfully",
		"data":    map[string]interface{}{},
	})
}
