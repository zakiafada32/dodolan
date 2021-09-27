package cart

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/api/common"
	"github.com/zakiafada32/retail/api/utils"
	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/cart"
)

type CartController struct {
	service cart.Service
}

func NewCartController(service cart.Service) *CartController {
	return &CartController{
		service: service,
	}
}

func (cont *CartController) AddCartItem(c echo.Context) error {
	var body cartRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	if err := c.Validate(&body); err != nil {
		return err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaimsUser)
	userId := claims.ID

	err := cont.service.UpdateCartItem(userId, body.convertToCartItemBusiness())
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.SucessCreated, echo.Map{}))
}
