package payment

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/api/common"
	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/payment"
)

type PaymentController struct {
	service payment.Service
}

func NewPaymentController(service payment.Service) *PaymentController {
	return &PaymentController{
		service: service,
	}
}

func (cont *PaymentController) FindAll(c echo.Context) error {
	payments, err := cont.service.FindAll()
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"payments": payments,
	}))
}

func (cont *PaymentController) CreateNew(c echo.Context) error {
	var body paymentProviderRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	err := cont.service.CreateNew(body.convertToPaymentProviderBusiness())
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.SucessCreated, echo.Map{}))
}

func (cont *PaymentController) Update(c echo.Context) error {
	var body paymentProviderRequestBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	id := c.Param("id")
	paymentId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	payment, err := cont.service.Update(uint32(paymentId), body.Name, body.Description)
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"payment": payment,
	}))
}
