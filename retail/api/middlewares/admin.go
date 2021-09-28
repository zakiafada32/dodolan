package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/api/utils"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*utils.JwtCustomClaimsUser)
		admin := claims.Admin
		if admin {
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}
