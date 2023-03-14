package util

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

func makeSlice(n int) []int {
	mySlice := make([]int, n)
	for i := 0; i < n; i++ {
		mySlice[i] = i
	}
	return mySlice
}

func testMany(n int, f func()) {
	for range makeSlice(100) {
		f()
	}
}

func TestRandomString(t *testing.T) {
	for _, i := range makeSlice(100) {
		res := RandomString(i)
		if i == 0 {
			require.Empty(t, res)
		} else {
			require.NotEmpty(t, res)
			require.Equal(t, i, len(res))
		}
	}
}

func TestRandomInt(t *testing.T) {
	testMany(100, func() {
		res := RandomInt(1, 100)
		require.NotEmpty(t, res)
		require.GreaterOrEqual(t, res, int64(1))
		require.LessOrEqual(t, res, int64(100))
	})
}

func TestRandomOwner(t *testing.T) {
	testMany(100, func() {
		res := RandomOwner()
		require.NotEmpty(t, res)
		require.Equal(t, 6, len(res))
		require.Subset(t, strings.Split(alphabet, ""), strings.Split(res, ""))
	})
}

func TestRandomMoney(t *testing.T) {
	testMany(100, func() {
		res := RandomMoney()
		require.NotEmpty(t, res)
		require.GreaterOrEqual(t, res, int64(0))
		require.LessOrEqual(t, res, int64(1000))
	})
}

func TestRandomAccountId(t *testing.T) {
	testMany(100, func() {
		res := RandomAccountId()
		require.NotEmpty(t, res)
		require.GreaterOrEqual(t, res, int64(0))
		require.LessOrEqual(t, res, int64(100000))
	})
}

func TestRandomCurrency(t *testing.T) {
	testMany(100, func() {
		res := RandomCurrency()
		require.NotEmpty(t, res)
		require.True(t, IsSupportedCurrency(res), fmt.Sprintf("%v should be supported currency", res))
	})
}

func TestRandomEmail(t *testing.T) {
	validate := validator.New()
	var err error
	testMany(100, func() {
		res := RandomEmail()
		require.NotEmpty(t, res)
		err = validate.Var(res, "required,email")
		require.NoError(t, err)
	})
}
