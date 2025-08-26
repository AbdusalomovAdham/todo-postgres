package controllers

import (
	"github.com/go-playground/validator/v10"
)

// validate user
var validate = validator.New()

func ValidateStruct(u interface{}) []map[string]string {
	err := validate.Struct(u)

	if errs, ok := err.(validator.ValidationErrors); ok {
		var validationErrors []map[string]string
		for _, e := range errs {
			validationErrors = append(validationErrors, map[string]string{
				"field":   e.Field(),
				"message": msgForTag(e.Tag()),
			})
		}
		return validationErrors
	}
	return nil
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "min":
		return "Too short"
	case "max":
		return "Too long"
	case "email":
		return "Invalid email format"
	}
	return "Invalid value"
}
