package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSupportedCurrency(t *testing.T) {
	for _, currency := range []string{USD, EUR, CAD, JPY} {
		require.True(t, IsSupportedCurrency(currency))
	}
}
func TestIsSupportedCurrencyNegative(t *testing.T) {
	testMany(100, func() {
		rand := RandomString(3)
		require.False(t, IsSupportedCurrency(rand))
	})
}
