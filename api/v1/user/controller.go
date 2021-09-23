package user

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/business/user"
)

type UserController struct {
	service user.Service
}

func NewUserController(service user.Service) *UserController {
	return &UserController{
		service: service,
	}
}

func (controller *UserController) CreateNewUser(c echo.Context) error {

	var body CreateNewUserRequestBody
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(&body); err != nil {
		log.Print(err)
		return err
	}

	err := controller.service.CreateNewUser(body.convertToUserBusiness())
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}
