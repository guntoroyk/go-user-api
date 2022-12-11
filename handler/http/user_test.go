package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/guntoroyk/go-user-api/entity"
	"github.com/guntoroyk/go-user-api/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	successGetUsersJSON   = "{\"code\":200,\"data\":[{\"id\":1,\"name\":\"User\",\"username\":\"User1\",\"phone\":\"\",\"role\":\"\",\"created_at\":\"2021-01-01 00:00:00\",\"updated_at\":\"2021-01-01 00:00:00\"}]}\n"
	successGetUserJSON    = "{\"code\":200,\"data\":{\"id\":1,\"name\":\"User\",\"username\":\"User1\",\"phone\":\"\",\"role\":\"\",\"created_at\":\"2021-01-01 00:00:00\",\"updated_at\":\"2021-01-01 00:00:00\"}}\n"
	successUpdateUserJSON = "{\"code\":200,\"data\":{\"id\":1,\"name\":\"User\",\"username\":\"User1\",\"phone\":\"\",\"role\":\"\",\"created_at\":\"2021-01-01 00:00:00\",\"updated_at\":\"2021-01-01 00:00:00\"}}\n"
	successCreateUserJSON = "{\"code\":201,\"data\":{\"id\":1,\"name\":\"User\",\"username\":\"User1\",\"phone\":\"\",\"role\":\"\",\"created_at\":\"2021-01-01 00:00:00\",\"updated_at\":\"2021-01-01 00:00:00\"}}\n"
)

func Test_handler_GetUsers(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	t.Run("success get users", func(t *testing.T) {
		mockTeamUC := mocks.NewMockUserUsecaseItf(ctrl)
		mockTeamUC.EXPECT().GetUsers().Return([]*entity.User{
			{
				ID:        1,
				Name:      "User",
				Username:  "User1",
				CreatedAt: "2021-01-01 00:00:00",
				UpdatedAt: "2021-01-01 00:00:00",
			},
		}, nil)
		h := NewHandler(mockTeamUC, nil)

		// Assertions
		if assert.NoError(t, h.GetUsers(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, successGetUsersJSON, rec.Body.String())
		}
	})
}

func Test_handler_GetUser(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")

	t.Run("success get user", func(t *testing.T) {
		mockTeamUC := mocks.NewMockUserUsecaseItf(ctrl)
		mockTeamUC.EXPECT().GetUser(gomock.Any()).Return(&entity.User{

			ID:        1,
			Name:      "User",
			Username:  "User1",
			CreatedAt: "2021-01-01 00:00:00",
			UpdatedAt: "2021-01-01 00:00:00",
		}, nil)
		h := NewHandler(mockTeamUC, nil)

		// Assertions
		if assert.NoError(t, h.GetUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, successGetUserJSON, rec.Body.String())
		}
	})
}

func Test_handler_UpdateUser(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")

	t.Run("success update user", func(t *testing.T) {
		mockTeamUC := mocks.NewMockUserUsecaseItf(ctrl)
		mockTeamUC.EXPECT().UpdateUser(gomock.Any()).Return(&entity.User{

			ID:        1,
			Name:      "User",
			Username:  "User1",
			CreatedAt: "2021-01-01 00:00:00",
			UpdatedAt: "2021-01-01 00:00:00",
		}, nil)
		h := NewHandler(mockTeamUC, nil)

		// Assertions
		if assert.NoError(t, h.UpdateUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, successUpdateUserJSON, rec.Body.String())
		}
	})
}

func Test_handler_CreateUser(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	t.Run("success create user", func(t *testing.T) {
		mockTeamUC := mocks.NewMockUserUsecaseItf(ctrl)
		mockTeamUC.EXPECT().CreateUser(gomock.Any()).Return(&entity.User{
			ID:        1,
			Name:      "User",
			Username:  "User1",
			CreatedAt: "2021-01-01 00:00:00",
			UpdatedAt: "2021-01-01 00:00:00",
		}, nil)
		h := NewHandler(mockTeamUC, nil)

		// Assertions
		if assert.NoError(t, h.CreateUser(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, successCreateUserJSON, rec.Body.String())
		}
	})
}

func Test_handler_DeleteUser(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")

	t.Run("success delete user", func(t *testing.T) {
		mockTeamUC := mocks.NewMockUserUsecaseItf(ctrl)
		mockTeamUC.EXPECT().DeleteUser(gomock.Any()).Return(nil)
		h := NewHandler(mockTeamUC, nil)

		// Assertions
		if assert.NoError(t, h.DeleteUser(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}
	})
}
