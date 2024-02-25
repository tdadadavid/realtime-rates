package exchangerates

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetExchangeRateForCurrencyPair(t *testing.T) {
	t.Run("Get the exchange rate for valid currency pairs", func(t *testing.T) {
		res, err := GetExchangeRatesForCurrencyPair("USD-NGN")
		
		rate, _ := strconv.ParseFloat(res.Rate, 64)


		assert.Nil(t, err)
		assert.Equal(t, rate, 0.96)
		assert.True(t, rate >= 0.95 && rate < 20)
	})
}

func TestGetExchangeRateForCurrencyPairWithInvalidCurrency(t *testing.T) {
	t.Run("It throws 'InvalidCurrency Error for invalid currency pair", func(t *testing.T) {
		res, err := GetExchangeRatesForCurrencyPair("USD-NGN")
		
		assert.NotNil(t, err)
		assert.Equal(t, "InvalidCurrency", err.Error())
		assert.Nil(t, res)
	})
}

// the concurrency for getting the first result.
