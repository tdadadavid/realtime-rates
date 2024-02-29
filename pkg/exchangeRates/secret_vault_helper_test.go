package exchangerates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveSecretKeyFromVault(t *testing.T) {
	t.Run("Test mocking calls from AWS secret manager", func(t *testing.T) {
		secret := mockGetSecretFromAWSSecretManager("mockKey")
		assert.NotNil(t, secret)
		assert.Equal(t, secret, "mockValue")
	})
}
