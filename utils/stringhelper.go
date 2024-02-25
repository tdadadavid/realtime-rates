package utils

import "strings"

type ExchnageRateCurrencies struct {
	From string 
	To string
}


func SplitCurrencyPair(currencyPair string) ExchnageRateCurrencies {
	currencies := strings.Split(currencyPair, "-")
	return formatCurrencies(currencies[0], currencies[1])
}


func formatCurrencies(firstCurrency, secondCurrency string) ExchnageRateCurrencies {
	return ExchnageRateCurrencies{
		From: firstCurrency,
		To: secondCurrency,
	}
}