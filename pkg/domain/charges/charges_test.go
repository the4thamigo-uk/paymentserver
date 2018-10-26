package charges

import (
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
	"testing"
)

func newTestCharges() Charges {
	return Charges{
		BearerCode: SHAR,
		Sender: map[string]money.Money{
			"USD": money.MustParse("234.56", "USD"),
			"GBP": money.MustParse("123.45", "GBP"),
		},
		Receiver: money.MustParse("234.56", "USD"),
	}
}

func TestCharges_Validate(t *testing.T) {
	c := newTestCharges()

	err := c.Validate("USD", "GBP")
	require.Nil(t, err)
}

func TestCharges_MissingSenderChargeSenderCcyError(t *testing.T) {
	c := newTestCharges()
	delete(c.Sender, "USD")

	err := c.Validate("USD", "GBP")
	require.NotNil(t, err.(ErrChargeNotFound))
}

func TestCharges_MissingSenderChargeReceiverCcyError(t *testing.T) {
	c := newTestCharges()
	delete(c.Sender, "GBP")

	err := c.Validate("USD", "GBP")
	require.NotNil(t, err.(ErrChargeNotFound))
}

func TestCharges_UnexpectedChargeInCcyError(t *testing.T) {
	c := newTestCharges()
	c.Sender["CLP"] = money.MustParse("123", "CLP")

	err := c.Validate("USD", "GBP")
	require.NotNil(t, err.(ErrChargeNotFound))
}

func TestCharges_SameCurrenciesTwoSenderChargesError(t *testing.T) {
	c := newTestCharges()

	err := c.Validate("USD", "USD")
	require.NotNil(t, err.(ErrChargeNotFound))
}
