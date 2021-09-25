package courier

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/business/courier"
)

type CourierController struct {
	service courier.Service
}

func NewCourierController(service courier.Service) *CourierController {
	return &CourierController{
		service: service,
	}
}

func (uc *CourierController) CreateNewCourierProvider(c echo.Context) error {

	var body createNewCourierProviderRequestBody
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	err := uc.service.CreateNewCourierProvider(body.convertToCourierProviderBusiness())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"code":    "new_courier_provider_created",
		"message": "new courier provider has been created successfully",
		"data":    map[string]interface{}{},
	})
}
