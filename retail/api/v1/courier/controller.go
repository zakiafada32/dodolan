package courier

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/api/common"
	"github.com/zakiafada32/retail/business"
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

func (cont *CourierController) FindAll(c echo.Context) error {
	couriers, err := cont.service.FindAll()
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"couriers": couriers,
	}))
}

func (cont *CourierController) CreateNew(c echo.Context) error {
	var body courierProviderRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	err := cont.service.CreateNew(body.convertToCourierProviderBusiness())
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.SuccessCreated, echo.Map{}))
}

func (cont *CourierController) Update(c echo.Context) error {
	var body courierProviderRequestBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	id := c.Param("id")
	courierId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	courier, err := cont.service.Update(uint32(courierId), body.Name, body.Description)
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"courier": courier,
	}))
}
