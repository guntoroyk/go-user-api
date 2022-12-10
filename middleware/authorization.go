package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/guntoroyk/go-user-api/entity"
	"github.com/labstack/echo/v4"
)

type middleware struct {
}

// NewMiddleware will create a new middleware
func NewMiddleware() *middleware {
	return &middleware{}
}

// Authorize is the middleware for authorization
func (m *middleware) Authorize(roles ...entity.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*entity.JwtCustomClaims)
			role := claims.Role

			for _, r := range roles {
				if role == string(r) {
					return next(c)
				}
			}

			return echo.ErrUnauthorized
		}
	}
}
