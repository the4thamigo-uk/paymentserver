package account

import (
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/bank"
	"testing"
)

func newTestAccount() Account {
	accType := 0
	return Account{
		ID: Identifier{
			Number: "GB82WEST12345698765432",
			Code:   IBAN},
		Name:    "John Smith",
		Address: "1 Test Street",
		Type:    &accType,
		BankID: bank.Identifier{
			ID:   "403000",
			Code: bank.GBDSC},
	}
}

func TestAccount_Validate(t *testing.T) {
	a := newTestAccount()
	err := a.Validate()
	require.Nil(t, err)
}

func TestAccount_ValidateNilType(t *testing.T) {
	a := newTestAccount()
	a.Type = nil
	err := a.Validate()
	require.Nil(t, err)
}

func TestAccount_ValidateNoNumberError(t *testing.T) {
	a := newTestAccount()
	a.ID.Number = ""
	err := a.Validate()
	require.NotNil(t, err.(ErrNotValid))
}

func TestAccount_ValidateNoNameError(t *testing.T) {
	a := newTestAccount()
	a.Name = ""
	err := a.Validate()
	require.NotNil(t, err.(ErrNotValid))
}

func TestAccount_ValidateNoAddressError(t *testing.T) {
	a := newTestAccount()
	a.Address = ""
	err := a.Validate()
	require.NotNil(t, err.(ErrNotValid))
}

func TestAccount_ValidateNoBankIdError(t *testing.T) {
	a := newTestAccount()
	a.BankID.ID = ""
	err := a.Validate()
	require.NotNil(t, err.(ErrNotValid))
}
