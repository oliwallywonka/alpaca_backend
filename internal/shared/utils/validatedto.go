package utils

import (
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(dto interface{}) []string {

	if err := validator.New().Struct(dto); err != nil {
		errors := make([]string, 0)
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, generateErrorString(err))
		}
		return errors
	}
	return nil
}

func generateErrorString(err validator.FieldError) string {
	filed, tag, param := err.Field(), err.Tag(), err.Param()
	switch tag {
	case "required":
		return filed + " the value is required"
	case "min":
		return filed + " the value must be greater than " + param
	case "max":
		return filed + " the value must be less than " + param
	case "oneof":
		return filed + " the value is must be one of asc or desc"
	default:
		return filed + " the value is invalid"
	}
}
