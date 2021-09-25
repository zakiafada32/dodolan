package user

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/api/common"
	"github.com/zakiafada32/retail/api/utils"
	"github.com/zakiafada32/retail/business"
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

func (uc *UserController) CreateNewUser(c echo.Context) error {

	var body createNewUserRequestBody
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	err := uc.service.CreateNewUser(body.convertToUserBusiness())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"code":    "new_user_created",
		"message": "new user created has been created successfully",
		"data":    map[string]interface{}{},
	})
}

func (uc *UserController) GetCurrentUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaimsUser)
	userId := claims.ID

	userData, err := uc.service.GetCurrentUser(userId)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	userResponse := convertToUserResponse(userData)
	return c.JSON(http.StatusOK, echo.Map{
		"user": userResponse,
	})

}

func (uc *UserController) Login(c echo.Context) error {
	var body loginRequestBody

	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	token, err := uc.service.Login(body.Email, body.Password)
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	return c.JSON(common.ConstructResponse(business.LoginSuccess, echo.Map{"token": token}))
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	var body updateUserRequestBody

	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaimsUser)
	userId := claims.ID

	userData, err := uc.service.UpdateUser(userId, body.convertToUpdateUserBusiness())
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	userResponse := convertToUserResponse(userData)
	return c.JSON(http.StatusOK, echo.Map{
		"user": userResponse,
	})
}
