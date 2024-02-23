package handlers

import (
	"net/http"
	exchangerates "realtime-exchange-rates/pkg/exchangeRates"

	"github.com/gofiber/fiber"
)

type CurrencyPairRequest struct {
	CurrencyPair string 
}

func HandleExchangeRequest(ctx *fiber.Ctx)  {
	body := new(CurrencyPairRequest)
	err := ctx.BodyParser(body)
  if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(err.Error())
		return;
	}

	exchangeRates, err := exchangerates.GetExchangeRatesForCurrencyPair(body)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(err.Error())
		return;
	}

	ctx.Status(http.StatusOK).JSON(exchangeRates)
}