package exchangerates

import (
	"fmt"
	"os"
	"realtime-exchange-rates/utils"
	"strings"
)

type ExchangeRateResult struct {
	Rate string `json:"rate"`
}

var (
	FrankFurterUrl       string = "https://api.frankfurter.app/latest?from=%s&to=%s"
	OpenExchangeRatesUrl        = "https://api.apilayer.com/fixer/latest?base=%s&symbols=%s"
)

func GetExchangeRatesForCurrencyPair(val string) (ExchangeRateResult, error) {
	currencies := utils.FormatCurrencies(val)

	urls := []string{prepareFixerUrl(currencies), prepareFrankFurterUrl(currencies)}
	result := make(chan string)

	for _, url := range urls {
		go GetExchangeRates(url, currencies, result)
	}

	return ExchangeRateResult{
		Rate: <-result,
	}, nil
}

// for unit testing purposes.
func GetExchangeRates(url string, currencies utils.ExchnageRateCurrencies, result chan string) {
	headers := utils.RequestParams{
		Url: url,
		Key: "",
	}

	// only request for api-keys when using fixer api
	// because only fixer requires api key to consume thier service.
	if strings.Contains(url, "fixer") {
		headers.Key = utils.GetSecretFromVault(os.Getenv("SECRET_NAME"))
	}

	response, err := utils.HandleRequest(headers)
	if err != nil {
		result <- ""
	}

	formatedResponse, err := FormatAPIResponse(response, url)
	if err != nil {
		result <- ""
	}

	rate := fmt.Sprintf("%f", formatedResponse[currencies.To])

	if err != nil {
		result <- ""
	} else {
		result <- rate
	}
}

func prepareFixerUrl(currencies utils.ExchnageRateCurrencies) string {
	return fmt.Sprintf(OpenExchangeRatesUrl, currencies.From, currencies.To)
}

func prepareFrankFurterUrl(currencies utils.ExchnageRateCurrencies) string {
	return fmt.Sprintf(FrankFurterUrl, currencies.From, currencies.To)
}
