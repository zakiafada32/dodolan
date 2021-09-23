package user

import (
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

func (u *UserController) CreateNewUser(c echo.Context) error {

	var body CreateNewUserRequestBody
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	err := u.service.CreateNewUser(body.convertToUserBusiness())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"code":    http.StatusCreated,
		"message": "new user created",
		"data":    map[string]interface{}{},
	})
}

func (u *UserController) Login(c echo.Context) error {
	var body LoginRequestBody

	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	token, err := u.service.Login(body.Email, body.Password)
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusOK,
		"message": "login success",
		"data": map[string]interface{}{
			"token": token,
		},
	})

}
