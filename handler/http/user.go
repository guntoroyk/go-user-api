package http

import (
	"errors"
	"net/http"

	"github.com/guntoroyk/go-user-api/entity"
	"github.com/guntoroyk/go-user-api/lib/converter"
	"github.com/labstack/echo/v4"
)

// GetUsers will get all users
func (h *handler) GetUsers(c echo.Context) error {
	users, err := h.userUsecase.GetUsers()

	resp := HttpResponse{
		Code: http.StatusOK,
		Data: users,
	}

	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	}

	return c.JSON(resp.Code, resp)
}

// GetUser will get user by id
func (h *handler) GetUser(c echo.Context) error {
	id := c.Param("id")

	user, err := h.userUsecase.GetUser(converter.ToInt(id))

	resp := HttpResponse{}

	if errors.Is(err, entity.ErrUserNotFound) {
		resp.Code = http.StatusNotFound
		resp.Error = err.Error()
	} else if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusOK
		resp.Data = user
	}

	return c.JSON(resp.Code, resp)
}

// CreateUser will create a user
func (h *handler) CreateUser(c echo.Context) error {
	user := new(entity.User)
	resp := HttpResponse{}

	if err := c.Bind(user); err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = err.Error()

		return c.JSON(resp.Code, resp)
	}

	user, err := h.userUsecase.CreateUser(user)

	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusCreated
		resp.Data = user
	}

	return c.JSON(resp.Code, resp)
}

// UpdateUser will update a user
func (h *handler) UpdateUser(c echo.Context) error {
	user := new(entity.User)
	resp := HttpResponse{}

	if err := c.Bind(user); err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = err.Error()

		return c.JSON(resp.Code, resp)
	}

	userID := c.Param("id")
	user.ID = converter.ToInt(userID)

	user, err := h.userUsecase.UpdateUser(user)

	if errors.Is(err, entity.ErrUserNotFound) {
		resp.Code = http.StatusNotFound
		resp.Error = err.Error()
	} else if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusOK
		resp.Data = user
	}

	return c.JSON(resp.Code, resp)
}

// DeleteUser will delete a user
func (h *handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := h.userUsecase.DeleteUser(converter.ToInt(id))

	resp := HttpResponse{}

	if errors.Is(err, entity.ErrUserNotFound) {
		resp.Code = http.StatusNotFound
		resp.Error = err.Error()
	} else if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusNoContent
	}

	return c.JSON(resp.Code, resp)
}
