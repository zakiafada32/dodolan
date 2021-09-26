package category

import (
	"net/http"
	"strconv"

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

func (cont *CategoryController) FindAllCategory(c echo.Context) error {
	categories, err := cont.service.FindAllCategory()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"category": categories,
	})
}

func (cont *CategoryController) FindCategoryById(c echo.Context) error {
	id := c.Param("id")
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	category, err := cont.service.FindCategoryById(uint32(categoryId))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"category": category,
	})
}

func (cont *CategoryController) CreateNewCategory(c echo.Context) error {
	var body createNewCategoryRequestBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	if err := c.Validate(&body); err != nil {
		return err
	}

	if err := cont.service.CreateNewCategory(body.convertToCategoryBusiness()); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"code":    "new_category_created",
		"message": "new category has been created successfully",
		"data":    echo.Map{},
	})
}

func (cont *CategoryController) UpdateCategory(c echo.Context) error {
	var body updateCategoryRequestBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	id := c.Param("id")
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	category, err := cont.service.UpdateCategory(uint32(categoryId), body.Name, body.Description)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"category": category,
	})
}
