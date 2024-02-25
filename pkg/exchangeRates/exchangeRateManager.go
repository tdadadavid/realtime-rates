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
		go getExchangeRates(url, result)
	}

	return ExchangeRateResult {
		Rate: <- result,
	}, nil
}

func getExchangeRates(url string, result chan string) {
	headers := utils.ApiKeyHeader {
		Url: url,
		Key: "sqt5JfnEkVOGaiTA63pA5EUyjPBiCzGA",
	}
	response, err := utils.HandleRequest(headers)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		result <- ""
	}else{
		result <- response
	}
}

func prepareFixerUrl(currencies utils.ExchnageRateCurrencies) string {
	return fmt.Sprintf(OpenExchangeRatesUrl, currencies.From, currencies.To)
}

func prepareFrankFurterUrl(currencies utils.ExchnageRateCurrencies) string {
	return fmt.Sprintf(FrankFurterUrl, currencies.From, currencies.To)
}