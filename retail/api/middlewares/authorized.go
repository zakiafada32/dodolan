package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zakiafada32/retail/api/utils"
)

func middlewareAuthorized() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &utils.JwtCustomClaimsUser{},
		SigningKey: []byte(os.Getenv("JWT_KEY")),
	}
	return middleware.JWTWithConfig(config)
}

var Authorized echo.MiddlewareFunc = middlewareAuthorized()
