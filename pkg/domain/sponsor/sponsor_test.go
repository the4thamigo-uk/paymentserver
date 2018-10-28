package sponsor

import (
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/bank"
	"testing"
)

func newTestSponsor() Sponsor {
	return Sponsor{
		Number: "123456789",
		BankID: bank.Identifier{
			ID:   "987654321",
			Code: bank.GBDSC,
		},
	}
}

func TestSponsor_Validate(t *testing.T) {
	s := newTestSponsor()
	err := s.Validate()
	require.Nil(t, err)
}

func TestSponsor_ReferenceError(t *testing.T) {
	s := newTestSponsor()
	s.Number = ""
	err := s.Validate()
	require.NotNil(t, err.(ErrNumber))
}

func TestSponsor_BankIDError(t *testing.T) {
	s := newTestSponsor()
	s.BankID.ID = ""
	err := s.Validate()
	require.NotNil(t, err.(ErrBankID))
}
