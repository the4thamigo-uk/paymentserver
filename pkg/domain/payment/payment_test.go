package payment

import (
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/account"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/bank"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/date"
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

func newPayment() Payment {
	return Payment{
		Credit:         money.MustParse("123.45", "USD"),
		Beneficiary:    newBeneficiary(),
		Debtor:         newDebtor(),
		ProcessingDate: date.MustParse("2000-02-01"),
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
