package api

import (
	"fmt"
	"net/http"
	"realtime-exchange-rates/pkg/persons"
	"realtime-exchange-rates/utils"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

type PersonQeuryParams struct {
	Sort string `query:"sort" validate:"oneof=asc desc"`
	GroupByCurrency string `query:"group_by_currency" validate:"oneof=true false"`
	USD string `query:"usd" validate:"numeric"`
}


func GetPersonsInformation(ctx *fiber.Ctx) error {

	queryValidator := validator.New()
	var queryParams PersonQeuryParams

	queryParams.Sort = ctx.Query("sort", "asc") //get the order to sort the person, else use ascending
	queryParams.GroupByCurrency = ctx.Query("group_by_currency", "false") //should the result be groupby salary.
	queryParams.USD = ctx.Query("usd", "10") // filter out users that have salaries less than 10 dollars

	if err := queryValidator.Struct(&queryParams); err != nil {
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


	filePath := "person.json"
	people, err := persons.GetPersons(filePath)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error retrieving persons",
			"error": err.Error(),
		})
	}

	// if user wants to get filtered persons.
	if queryParams.USD != "" {
		amountConstraint, err := strconv.ParseFloat(queryParams.USD, 64);

		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Error retrieving persons",
				"error": err.Error(),
			})
		}

		result, err := people.FilterBySalary(float64(amountConstraint))
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Error filtering people",
				"error": err.Error(),
			})
		}

		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": fmt.Sprintf("People with salary greater/equal to %f USD", amountConstraint),
			"data": result.Data,
		});

	}

	// retrieve people grouped by currency.
	if strings.ToLower(queryParams.GroupByCurrency) == "true" {
		results := people.GroupByCurrency()

		return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Persons.",
		"data": results,
	});
	}

	// by default sort people.
	result := people.Sort(queryParams.Sort).Data

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Persons.",
		"data": result,
	});
}