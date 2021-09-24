package product

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

func (pc *ProductController) CreateNewProduct(c echo.Context) error {
	var body createNewProductRequestBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	if err := c.Validate(&body); err != nil {
		return err
	}

	if err := pc.service.CreateNewProduct(body.convertToProductBusiness()); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"code":    "new_product_created",
		"message": "new product has been created successfully",
		"data":    map[string]interface{}{},
	})
}
