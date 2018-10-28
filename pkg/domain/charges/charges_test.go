package charges

import (
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
	"testing"
)

func newTestCharges() Charges {
	return Charges{
		BearerCode: SHAR,
		Sender: []money.Money{
			money.MustParse("234.56", "USD"),
			money.MustParse("123.45", "GBP"),
		},
		Receiver: money.MustParse("234.56", "USD"),
	}
}

func TestCharges_Validate(t *testing.T) {
	c := newTestCharges()

	err := c.Validate("USD", "GBP")
	require.Nil(t, err)
}

func TestCharges_MissingSenderChargeUSDError(t *testing.T) {
	c := newTestCharges()
	c.Sender = c.Sender[0:1]
	err := c.Validate("USD", "GBP")
	require.NotNil(t, err.(ErrChargeNotFound))
}

func TestCharges_MissingSenderChargeGBPError(t *testing.T) {
	c := newTestCharges()
	c.Sender = c.Sender[1:]

	err := c.Validate("USD", "GBP")
	require.NotNil(t, err.(ErrChargeNotFound))
}

func TestCharges_MissingReceiverChargeWrongCcyError(t *testing.T) {
	c := newTestCharges()
	c.Receiver = money.MustParse("123.45", "GBP")

	err := c.Validate("USD", "GBP")
	require.NotNil(t, err.(ErrChargeNotFound))
}

func TestCharges_UnexpectedChargeInCcyError(t *testing.T) {
	c := newTestCharges()
	c.Sender = append(c.Sender, money.MustParse("123", "CLP"))

	err := c.Validate("USD", "GBP")
	require.NotNil(t, err.(ErrChargeNotFound))
}

func TestCharges_SameCurrenciesTwoSenderChargesError(t *testing.T) {
	c := newTestCharges()

	err := c.Validate("USD", "USD")
	require.NotNil(t, err.(ErrChargeNotFound))
}
