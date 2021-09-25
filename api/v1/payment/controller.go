package payment

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

func (pc *PaymentController) CreateNewPaymentProvider(c echo.Context) error {
	var body createNewPaymentProviderRequestBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	if err := c.Validate(&body); err != nil {
		return err
	}

	if err := pc.service.CreateNewPaymentProvider(body.convertToPaymentProviderBusiness()); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"code":    "new_product_created",
		"message": "new payment has been created successfully",
		"data":    map[string]interface{}{},
	})
}
