package order

import (
	"github.com/labstack/echo/v4"
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
	return nil
}

func (cont *OrderController) Payment(c echo.Context) error {
	return nil
}

func (cont *OrderController) Courier(c echo.Context) error {
	return nil
}
