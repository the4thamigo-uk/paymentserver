package payment

import (
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/account"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/amount"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/bank"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/charges"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/currency"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/date"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/fx"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
	"testing"
)

func newBeneficiary() account.Account {
	accType := 0
	return account.Account{
		ID: account.Identifier{
			Number: "GB82WEST12345698765432",
			Code:   account.IBAN},
		Name:    "John Smith",
		Address: "1 Test Street",
		Type:    &accType,
		BankID: bank.Identifier{
			ID:   "403000",
			Code: bank.GBDSC},
	}
}

func newDebtor() account.Account {
	accType := 0
	return account.Account{
		ID: account.Identifier{
			Number: "12345678",
			Code:   account.BBAN},
		Name:    "John Doe",
		Address: "2 Test Street",
		Type:    &accType,
		BankID: bank.Identifier{
			ID:   "404000",
			Code: bank.GBDSC},
	}
}

func newTestCharges() charges.Charges {
	return charges.Charges{
		BearerCode: charges.SHAR,
		Sender: map[currency.Currency]money.Money{
			"USD": money.MustParse("234.56", "USD"),
			"GBP": money.MustParse("123.45", "GBP"),
		},
		Receiver: money.MustParse("234.56", "USD"),
	}
}

func newFx() *fx.Contract {
	return &fx.Contract{
		Reference: "FX123",
		Rate:      amount.MustParse("1.23456789"),
		Domestic:  money.MustParse("2.89", "USD"),
	}
}

func newPayment() Payment {
	return Payment{
		Credit:         money.MustParse("2.34", "GBP"),
		Beneficiary:    newBeneficiary(),
		Debtor:         newDebtor(),
		ProcessingDate: date.MustParse("2000-02-01"),
		Charges:        newTestCharges(),
		Fx:             newFx(),
	}
}

func TestPayment_Validate(t *testing.T) {
	p := newPayment()
	err := p.Validate()
	require.Nil(t, err)
}

func TestPayment_BeneficiaryError(t *testing.T) {
	p := newPayment()
	p.Beneficiary.Name = ""
	err := p.Validate()
	require.NotNil(t, err.(ErrBeneficiaryNotValid))
}

func TestPayment_DebtorError(t *testing.T) {
	p := newPayment()
	p.Debtor.Name = ""
	err := p.Validate()
	require.NotNil(t, err.(ErrDebtorNotValid))
}

func TestPayment_ChargesError(t *testing.T) {
	p := newPayment()
	delete(p.Charges.Sender, "USD")
	err := p.Validate()
	require.NotNil(t, err.(ErrChargesNotValid))
}

func TestPayment_NoFx(t *testing.T) {
	p := newPayment()
	p.Fx = nil
	err := p.Validate()
	require.Nil(t, err)
}

func TestPayment_FxError(t *testing.T) {
	p := newPayment()
	p.Fx.Reference = ""
	err := p.Validate()
	require.NotNil(t, err.(ErrChargesNotValid))
}
