package utils

import "github.com/go-playground/validator/v10"

var Validator = validator.New(validator.WithRequiredStructEnabled())

func ValidateRequestBody(userInput interface{}) error {
	return handleValidation(userInput)
}

func handleValidation(userInput interface{}) error {

	if err := Validator.Struct(userInput); err != nil {
		return err;
	}

	return nil
}