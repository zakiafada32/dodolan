package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/api/utils"
	"github.com/zakiafada32/retail/api/v1/user"
)

func RegisterPath(e *echo.Echo, userController *user.UserController) {
	if userController == nil {
		panic("user controller cannot be nil")
	}

	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	userV1 := e.Group("api/v1/users")
	userV1.POST("", userController.CreateNewUser)
}
