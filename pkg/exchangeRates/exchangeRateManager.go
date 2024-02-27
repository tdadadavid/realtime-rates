package exchangerates

import (
	"fmt"
	"realtime-exchange-rates/utils"
)

type ExchangeRateResult struct {
	Rate string `json:"rate"`
}



var (
	FrankFurterUrl  string = "https://api.frankfurter.app/latest?from=%s&to=%s"
	OpenExchangeRatesUrl = "https://api.apilayer.com/fixer/latest?base=%s&symbols=%s"
)

func GetExchangeRatesForCurrencyPair(val string) (ExchangeRateResult, error) {
	currencies := utils.SplitCurrencyPair(val)

	urls := []string{prepareFixerUrl(currencies), prepareFrankFurterUrl(currencies)}
	result := make(chan string)
	
	for _, url := range urls {
		go GetExchangeRates(url, currencies, result)
	}

	return ExchangeRateResult {
		Rate: <- result,
	}, nil
}

// for unit testing purposes.
func GetExchangeRates(url string, currencies utils.ExchnageRateCurrencies, result chan string) {
	headers := utils.RequestParams {
		Url: url,
		Key: "sqt5JfnEkVOGaiTA63pA5EUyjPBiCzGA",
	}

	response, err := utils.HandleRequest(headers)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		result <- ""
	}

	formatedResponse, err := FormatAPIResponse(response, url)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		result <- ""
	}

	rate := fmt.Sprintf("%f", formatedResponse[currencies.To])

	if err != nil {
		fmt.Println("Error: ", err.Error())
		result <- ""
	}else{
		result <- rate
	}
}

func prepareFixerUrl(currencies utils.ExchnageRateCurrencies) string {
	return fmt.Sprintf(OpenExchangeRatesUrl, currencies.From, currencies.To)
}

func prepareFrankFurterUrl(currencies utils.ExchnageRateCurrencies) string {
	return fmt.Sprintf(FrankFurterUrl, currencies.From, currencies.To)
}