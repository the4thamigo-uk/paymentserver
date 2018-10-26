package payment

import (
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/account"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/date"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
)

// Payment defines a payment from a debtor to a beneficiary
type Payment struct {
	Credit         money.Money
	Beneficiary    account.Account
	Debtor         account.Account
	ProcessingDate date.Date
}

// Validate performs some basic checks on the validity of the Payment
func (p *Payment) Validate() error {
	err := p.Beneficiary.Validate()
	if err != nil {
		return errBeneficiaryNotValid(err)
	}
	err = p.Debtor.Validate()
	if err != nil {
		return errDebtorNotValid(err)
	}

	return nil
}
