package product

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/api/common"
	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/product"
)

type ProductController struct {
	service product.Service
}

func NewProductController(service product.Service) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (cont *ProductController) FindById(c echo.Context) error {
	id := c.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	product, err := cont.service.FindById(uint32(productId))
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"product": product,
	}))
}

func (cont *ProductController) FindAll(c echo.Context) error {

	products, err := cont.service.FindAll()
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"products": products,
	}))
}

func (cont *ProductController) CreateNew(c echo.Context) error {
	var body createNewProductRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	if err := c.Validate(&body); err != nil {
		return err
	}

	if err := cont.service.CreateNew(body.convertToProductBusiness()); err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.SucessCreated, echo.Map{}))
}

func (cont *ProductController) Update(c echo.Context) error {
	var body updateProductRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	id := c.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	product, err := cont.service.Update(uint32(productId), body.convertToUpdateProductBusiness())
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.SucessCreated, echo.Map{
		"product": product,
	}))
}
