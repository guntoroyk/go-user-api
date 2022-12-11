package http

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/guntoroyk/go-user-api/entity"
	"github.com/labstack/echo/v4"
)

// Login will login a user
func (h *handler) Login(c echo.Context) error {
	resp := HttpResponse{}
	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = err.Error()

		return c.JSON(resp.Code, resp)
	}

	token, err := h.authUsecase.Login(user.Username, user.Password)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusOK
		resp.Data = token
	}

	return c.JSON(resp.Code, resp)
}

// RefreshToken will refresh a token
func (h *handler) RefreshToken(c echo.Context) error {
	resp := HttpResponse{}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*entity.JwtCustomClaims)
	if claims.Type != entity.RefreshToken {
		resp.Code = http.StatusUnauthorized
		resp.Error = "Invalid token type"

		return c.JSON(resp.Code, resp)
	}

	newToken, err := h.authUsecase.RefreshToken(claims.Username, claims.Role)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusOK
		resp.Data = newToken
	}

	return c.JSON(resp.Code, resp)
}
