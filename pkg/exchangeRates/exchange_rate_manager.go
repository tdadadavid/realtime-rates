package exchangerates

import (
	"fmt"
	"log"
	"os"
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
	currencies := FormatCurrencies(val)

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
func GetExchangeRates(url string, currencies ExchnageRateCurrencies, result chan string) {
	headers := RequestParams{
		Url: url,
		Key: "",
	}

	// only request for api-keys when using fixer api
	// because only fixer requires api key to consume thier service.
	if strings.Contains(url, "fixer") {
		headers.Key = GetSecretFromVault(os.Getenv("SECRET_NAME"))
	}

	response, err := HandleRequest(headers)
	if err != nil {
		log.Println("[Request Error]: ", err)
		result <- ""
	}

	formatedResponse, err := FormatAPIResponse(response, url)
	if err != nil {
		log.Println("[Currency Format Error]: ", err)
		result <- ""
	}

	rate := fmt.Sprintf("%f", formatedResponse[currencies.To])

	if err != nil {
		log.Println("[Currency Format Error]: ", err)
		result <- ""
	} else {
		result <- rate
	}
}

func prepareFixerUrl(currencies ExchnageRateCurrencies) string {
	return fmt.Sprintf(OpenExchangeRatesUrl, currencies.From, currencies.To)
}

func prepareFrankFurterUrl(currencies ExchnageRateCurrencies) string {
	return fmt.Sprintf(FrankFurterUrl, currencies.From, currencies.To)
}
