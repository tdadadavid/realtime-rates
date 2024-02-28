package exchangerates

import (
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

		assert.NotNil(t, rate)
	})
}