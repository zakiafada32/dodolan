package order

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/api/common"
	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/order"
)

type OrderController struct {
	service order.Service
}

func NewOrderController(service order.Service) *OrderController {
	return &OrderController{
		service: service,
	}
}

func (cont *OrderController) FindAll(c echo.Context) error {
	return nil
}

func (cont *OrderController) FindById(c echo.Context) error {
	id := c.Param("id")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}
	order, err := cont.service.FindById(uint32(orderId))
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}
	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"order": order,
	}))
}

func (cont *OrderController) Payment(c echo.Context) error {
	return nil
}

func (cont *OrderController) Courier(c echo.Context) error {
	return nil
}
