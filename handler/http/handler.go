package http

import "github.com/guntoroyk/go-user-api/usecase"

type handler struct {
	userUsecase usecase.UserUsecaseItf
	authUsecase usecase.AuthUsecaseItf
}

// NewHandler is a constructor for handler
func NewHandler(
	userUsecase usecase.UserUsecaseItf,
	authUsecase usecase.AuthUsecaseItf,
) *handler {
	return &handler{
		userUsecase: userUsecase,
		authUsecase: authUsecase,
	}
}
