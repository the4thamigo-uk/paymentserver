package payment

import (
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/account"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/charges"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/date"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/fx"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
)

// Payment defines a payment from a debtor to a beneficiary
type Payment struct {
	Credit         money.Money
	Beneficiary    account.Account
	Debtor         account.Account
	Charges        charges.Charges
	Fx             *fx.Contract
	ID             string
	Purpose        string
	Type           string // TODO enum?
	Scheme         string // TODO are certain scheme combinations valid?
	SchemeType     string
	SchemeSubType  string
	EndToEndRef    string
	NumericRef     string // validate is numeric
	Reference      string
	ProcessingDate date.Date
	//Sponsor //TODO
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
	err = p.Charges.Validate("USD", p.Credit.Currency())
	if err != nil {
		return errChargesNotValid(err)
	}
	if p.Fx != nil {
		err = p.Fx.Validate(p.Credit)
		if err != nil {
			return errFxNotValid(err)
		}
	}
	return nil
}
