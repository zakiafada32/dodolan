package order

import (
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/api/common"
	"github.com/zakiafada32/retail/api/utils"
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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaimsUser)
	userId := claims.ID

	orders, err := cont.service.FindAll(userId)
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}
	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"orders": orders,
	}))
}

func (cont *OrderController) FindById(c echo.Context) error {
	id := c.Param("id")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaimsUser)
	userId := claims.ID

	order, err := cont.service.FindById(userId, uint32(orderId))
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}
	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"order": order,
	}))
}

func (cont *OrderController) Payment(c echo.Context) error {
	var body paymentRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaimsUser)
	userId := claims.ID

	if err := cont.service.Payment(userId, body.OrderId, body.TotalAmount); err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.Success, echo.Map{}))
}

func (cont *OrderController) Courier(c echo.Context) error {
	var body courierRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaimsUser)
	userId := claims.ID

	if err := cont.service.Courier(userId, body.OrderId); err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.Success, echo.Map{}))
}
