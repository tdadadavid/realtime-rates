package api

import (
	"fmt"
	"net/http"
	exchangerates "realtime-exchange-rates/pkg/exchangeRates"
	"realtime-exchange-rates/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


type CurrencyPairRequest struct {
	CurrencyPair string `json:"currency-pair" validate:"required"`
}

type CurrenPairResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data CurrencyPairExchangeRates `json:"data"`
}

type CurrencyPairExchangeRates map[string]string


func HandleRealtimeExchangeRate(ctx *fiber.Ctx) error {
	requestValidator := validator.New()
	var request CurrencyPairRequest
	// get the request
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body.",
			"errors": err.Error(),
		})
	}

	
	// // validate the request body
	if err := requestValidator.Struct(&request); err != nil {
		var errorMessages []string
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			for _, fieldError := range validationErrors {
				errMsg := utils.FormatErrorMessages(fieldError)
				errorMessages = append(errorMessages, errMsg)
			}
		}
		
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Validation errors.",
			"errors": errorMessages,
		})
	}

	// split the strings and get the two currency parties get the exchange rate
	result, err := exchangerates.GetExchangeRatesForCurrencyPair(request.CurrencyPair)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	// return response.
	return ctx.Status(http.StatusOK).JSON(CurrenPairResponse{
		Success: true,
		Message: fmt.Sprintf("Exchange rate for %s", request.CurrencyPair),
		Data: CurrencyPairExchangeRates{
			request.CurrencyPair: result.Rate,
		},
	})
}