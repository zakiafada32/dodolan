package category

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/api/common"
	"github.com/zakiafada32/retail/business"
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
	categories, err := cont.service.FindAll()
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"categories": categories,
	}))
}

func (cont *CategoryController) FindCategoryById(c echo.Context) error {
	id := c.Param("id")
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	category, err := cont.service.FindById(uint32(categoryId))
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"category": category,
	}))
}

func (cont *CategoryController) CreateNewCategory(c echo.Context) error {
	var body categoryRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	if err := c.Validate(&body); err != nil {
		return err
	}

	if err := cont.service.CreateNew(body.convertToCategoryBusiness()); err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.Success, echo.Map{}))
}

func (cont *CategoryController) UpdateCategory(c echo.Context) error {
	var body categoryRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	id := c.Param("id")
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	category, err := cont.service.Update(uint32(categoryId), body.Name, body.Description)
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"category": category,
	}))
}
