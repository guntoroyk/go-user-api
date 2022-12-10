package entity

import "errors"

var (
	// ErrUserNotFound is the error message when user is not found
	ErrUserNotFound = errors.New("user not found")
	// ErrWrongCredentials is the error message when credentials are wrong
	ErrWrongCredentials = errors.New("wrong credentials")
	// ErrUserAlreadyExist is the error message when user already exist
	ErrUserAlreadyExist = errors.New("user already exist")
)
