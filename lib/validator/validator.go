package validator

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

// GetValidator will return a validator instance
func GetValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}
	return validate
}
