package validator_helper

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func SetValidate() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func GetValidate() *validator.Validate {
	return validate
}
