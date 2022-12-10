package proxy

import (
	"log"

	"github.com/guntoroyk/go-user-api/entity"
	"github.com/guntoroyk/go-user-api/lib/crypto"
	"github.com/guntoroyk/go-user-api/repository"
)

type userRepoProxy struct {
	Repo repository.UserRepoItf
}

// NewUserRepoProxy will create new an UserRepo object representation of UserRepoItf interface
func NewUserRepoProxy(repo repository.UserRepoItf) *userRepoProxy {
	return &userRepoProxy{
		Repo: repo,
	}
}

// CreateUser will create new user
func (c *userRepoProxy) CreateUser(user *entity.User) (*entity.User, error) {
	// hash password
	hashedPasswd, err := crypto.HashPassword(user.Password)
	if err != nil {
		log.Println("func CreateUser error hashing password, ", err.Error())
		return nil, err
	}

	user.Password = hashedPasswd
	return c.Repo.CreateUser(user)
}

// GetUser will get user by id
func (c *userRepoProxy) GetUser(filter *entity.UserFilter) (*entity.User, error) {
	return c.Repo.GetUser(filter)
}

// GetUsers will get all users
func (c *userRepoProxy) GetUsers() ([]*entity.User, error) {
	return c.Repo.GetUsers()
}

// UpdateUser will update user by id
func (c *userRepoProxy) UpdateUser(user *entity.User) (*entity.User, error) {
	// hash password
	hashedPasswd, err := crypto.HashPassword(user.Password)
	if err != nil {
		log.Println("func UpdateUser error hashing password, ", err.Error())
		return nil, err
	}

	user.Password = hashedPasswd
	return c.Repo.UpdateUser(user)
}

// DeleteUser will delete user by id
func (c *userRepoProxy) DeleteUser(id int) error {
	return c.Repo.DeleteUser(id)
}
