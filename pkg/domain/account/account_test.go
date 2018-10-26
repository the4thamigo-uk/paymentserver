package account

import (
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/bank"
	"testing"
)

func newTest() Account {
	accType := 0
	return Account{
		Id: Identifier{
			Number: "GB82WEST12345698765432",
			Code:   IBAN},
		Name:    "John Smith",
		Address: "1 Test Street",
		Type:    &accType,
		bankId: bank.Identifier{
			Id:   "403000",
			Code: bank.GBDSC},
	}
}

func Test_Validate(t *testing.T) {
	a := newTest()
	err := a.Validate()
	require.Nil(t, err)
}

func Test_ValidateNilType(t *testing.T) {
	a := newTest()
	a.Type = nil
	err := a.Validate()
	require.Nil(t, err)
}

func Test_ValidateNoNumberError(t *testing.T) {
	a := newTest()
	a.Id.Number = ""
	err := a.Validate()
	require.NotNil(t, err.(ErrNotValid))
}

func Test_ValidateNoNameError(t *testing.T) {
	a := newTest()
	a.Name = ""
	err := a.Validate()
	require.NotNil(t, err.(ErrNotValid))
}

func Test_ValidateNoAddressError(t *testing.T) {
	a := newTest()
	a.Address = ""
	err := a.Validate()
	require.NotNil(t, err.(ErrNotValid))
}

func Test_ValidateNoBankIdError(t *testing.T) {
	a := newTest()
	a.bankId.Id = ""
	err := a.Validate()
	require.NotNil(t, err.(ErrNotValid))
}
