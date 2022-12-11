package usecase

import "github.com/guntoroyk/go-user-api/entity"

// AuthUsecaseItf is the interface for the AuthUsecase struct
type AuthUsecaseItf interface {
	Login(username, password string) (*entity.Token, error)
	RefreshToken(username, role string) (*entity.Token, error)
}

// UserUsecaseItf is the interface for the UserUsecase struct
type UserUsecaseItf interface {
	GetUsers() ([]*entity.User, error)
	GetUser(id int) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id int) error
}
