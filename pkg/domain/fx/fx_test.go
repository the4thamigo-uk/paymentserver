package fx

import (
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/amount"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
	"testing"
)

func newTestContract() (money.Money, Contract) {
	foreign := money.MustParse("2.34", "GBP")
	return foreign, Contract{
		Reference: "FX123",
		Rate:      amount.MustParse("1.23456789"),
		Domestic:  money.MustParse("2.89", "USD"),
	}
}

func TestContract_Validate(t *testing.T) {
	foreign, c := newTestContract()
	err := c.Validate(foreign)
	require.Nil(t, err)
}

func TestContract_ReferenceError(t *testing.T) {
	foreign, c := newTestContract()
	c.Reference = ""
	err := c.Validate(foreign)
	require.NotNil(t, err.(ErrReferenceNotValid))
}

func TestContract_ExchangeError(t *testing.T) {
	foreign, c := newTestContract()
	c.Rate = amount.Amount{}
	err := c.Validate(foreign)
	require.NotNil(t, err.(ErrExchange))
}

func TestContract_RateError(t *testing.T) {
	_, c := newTestContract()
	foreign := money.MustParse("2.35", "GBP")
	err := c.Validate(foreign)
	require.NotNil(t, err.(ErrRateNotValid))
}
