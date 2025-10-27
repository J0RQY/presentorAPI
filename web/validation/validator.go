package validation

import (
	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) interface{} {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, e := range validationErrors {
			errors[e.Field()] = GetErrorMessage(e)
		}
		return errors
	}
	return err.Error()
}

func GetErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Must be at least " + err.Param() + " characters"
	case "max":
		return "Must be at most " + err.Param() + " characters"
	default:
		return "Invalid value"
	}
}
