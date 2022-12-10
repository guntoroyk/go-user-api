package usecase

import (
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/guntoroyk/go-user-api/entity"
	"github.com/guntoroyk/go-user-api/lib/crypto"
	"github.com/guntoroyk/go-user-api/lib/jwt"
	"github.com/guntoroyk/go-user-api/repository"
)

var (
	tokenExpire        = time.Now().Add(time.Hour * 24)
	refreshTokenExpire = time.Now().Add(time.Hour * 24 * 7)
)

type authUsecase struct {
	userRepo     repository.UserRepoItf
	tokenService jwt.TokenServiceItf
	validator    *validator.Validate
}

// NewAuthUsecase will create new an AuthUsecase object representation of AuthUsecaseItf interface
func NewAuthUsecase(
	userRepo repository.UserRepoItf,
	tokenService jwt.TokenServiceItf,
	validator *validator.Validate,
) AuthUsecaseItf {
	return &authUsecase{
		userRepo:     userRepo,
		tokenService: tokenService,
		validator:    validator,
	}
}

// Login will login a user
func (c *authUsecase) Login(username, password string) (*entity.Token, error) {
	user, err := c.userRepo.GetUser(&entity.UserFilter{Username: username})
	if err != nil && err == sql.ErrNoRows {
		return nil, entity.ErrWrongCredentials
	}

	if !crypto.CheckPasswordHash(password, user.Password) {
		return nil, entity.ErrWrongCredentials
	}

	token, err := c.tokenService.GenerateToken(user.Username, string(user.Role), tokenExpire)
	if err != nil {
		return nil, err
	}

	refreshToken, err := c.tokenService.GenerateToken(user.Username, string(user.Role), refreshTokenExpire)
	if err != nil {
		return nil, err
	}

	return &entity.Token{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
