package sponsor

import (
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/bank"
)

// Sponsor specifies the details of the payment sponsor
type Sponsor struct {
	Number string
	BankID bank.Identifier
}

// Validate performs some checks on the Sponsor.
func (s *Sponsor) Validate() error {
	if s.Number == "" {
		return errNumber(s.Number)
	}
	err := s.BankID.Validate()
	if err != nil {
		return errBankID(err)
	}
	return nil
}
