package http

import (
	"net/http"

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
