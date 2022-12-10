package repository

import "github.com/guntoroyk/go-user-api/entity"

// UserRepoItf is the interface for the UserRepo struct
type UserRepoItf interface {
	GetUsers() ([]*entity.User, error)
	GetUser(filter *entity.UserFilter) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id int) error
}
