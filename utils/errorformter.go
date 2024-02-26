package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatErrorMessages(e validator.FieldError) string {
	fieldName := e.Field()
	tag := e.Tag()
	param := e.Param()

	var message string

	switch tag {
	case "required":
		message = fmt.Sprintf("Field %s is required", fieldName)
	case "oneof":
		message = fmt.Sprintf("%s field should be one of these (%s)", fieldName, param)
	case "eq":
		message = fmt.Sprintf("%s field should be equal to one of these (%s)", fieldName, param)
	case "numeric":
		message = fmt.Sprintf("%s field must be a number", fieldName)
	}

	return message
}