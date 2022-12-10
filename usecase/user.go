package usecase

import (
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/guntoroyk/go-user-api/entity"
	"github.com/guntoroyk/go-user-api/repository"
)

type userUsecase struct {
	userRepo  repository.UserRepoItf
	validator *validator.Validate
}

// NewUserUsecase will create new an UserUsecase object representation of UserUsecaseItf interface
func NewUserUsecase(userRepo repository.UserRepoItf, validator *validator.Validate) UserUsecaseItf {
	return &userUsecase{
		userRepo:  userRepo,
		validator: validator,
	}
}

// GetUsers will get all users
func (c *userUsecase) GetUsers() ([]*entity.User, error) {
	return c.userRepo.GetUsers()
}

// GetUser will get user by id
func (c *userUsecase) GetUser(id int) (*entity.User, error) {
	user, err := c.userRepo.GetUser(&entity.UserFilter{ID: id})
	if err != nil && err == sql.ErrNoRows {
		return nil, entity.ErrUserNotFound
	}
	return user, err
}

// CreateUser will create a user
func (c *userUsecase) CreateUser(user *entity.User) (*entity.User, error) {
	err := c.validator.Struct(user)
	if err != nil {
		return nil, err
	}

	existingUser, err := c.userRepo.GetUser(&entity.UserFilter{Username: user.Username})
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if existingUser != nil {
		return nil, entity.ErrUserAlreadyExist
	}

	return c.userRepo.CreateUser(user)
}

// UpdateUser will update a user
func (c *userUsecase) UpdateUser(user *entity.User) (*entity.User, error) {
	err := c.validator.Struct(user)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	user, err = c.userRepo.UpdateUser(user)
	if err != nil && err == sql.ErrNoRows {
		return nil, entity.ErrUserNotFound
	}
	return user, err
}

// DeleteUser will delete a user
func (c *userUsecase) DeleteUser(id int) error {
	err := c.userRepo.DeleteUser(id)
	if err != nil && err == sql.ErrNoRows {
		return entity.ErrUserNotFound
	}
	return err
}
