package utils

import "strings"

type ExchnageRateCurrencies struct {
	From string
	To   string
}

func splitCurrencyPair(currencyPair string) []string {
	return strings.Split(currencyPair, "-")
}

func FormatCurrencies(currenyPair string) ExchnageRateCurrencies {
	currencies := splitCurrencyPair(currenyPair)
	return ExchnageRateCurrencies{
		From: currencies[0],
		To:   currencies[1],
	}
}
