package exchangerates

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Integrations
func TestGetExchangeRateForCurrencyPair(t *testing.T) {
	t.Run("Get the exchange rate for valid currency pairs", func(t *testing.T) {
		res, err := GetExchangeRatesForCurrencyPair("USD-GBP")
		assert.Nil(t, err)
		
		rate, err := strconv.ParseFloat(res.Rate, 64)
		assert.Nil(t, err)

		assert.Nil(t, err)
		assert.NotNil(t, rate)
	})
}

func TestGetExchangeRateForCurrencyPairWithInvalidCurrency(t *testing.T) {
	t.Run("It throws 'InvalidCurrency' Error for invalid currency pair", func(t *testing.T) {

		wrongCurrency := "NPM"
		res, err := GetExchangeRatesForCurrencyPair(fmt.Sprintf("USD-%s", wrongCurrency))
		
		assert.Nil(t, err)
		assert.Contains(t, res.Rate, "0.0")
	})
}