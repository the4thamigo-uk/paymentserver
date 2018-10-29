package payment

import (
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/account"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/charges"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/date"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/fx"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/sponsor"
)

// Payment defines a payment from a debtor to a beneficiary
type Payment struct {
	OrganisationID string          `json:"organisation_id"`
	Credit         money.Money     `json:"credit"`
	Beneficiary    account.Account `json:"beneficiary"`
	Debtor         account.Account `json:"debtor"`
	Charges        charges.Charges `json:"charges"`
	Fx             *fx.Contract    `json:"fx"`
	ID             string          `json:"id"`
	Purpose        string          `json:"purpose"`
	Type           string          `json:"type"`   // TODO enum?
	Scheme         string          `json:"scheme"` // TODO are certain scheme combinations valid?
	SchemeType     string          `json:"scheme_type"`
	SchemeSubType  string          `json:"scheme_sub_type"`
	EndToEndRef    string          `json:"end_to_end_ref"`
	NumericRef     string          `json:"numeric_ref"` // validate is numeric
	Reference      string          `json:"ref"`
	ProcessingDate date.Date       `json:"processing_date"`
	Sponsor        sponsor.Sponsor `json:"sponsor"`
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
	err = p.Sponsor.Validate()
	if err != nil {
		return errSponsorNotValid(err)
	}

	fields := map[string]interface{}{
		"id":                p.ID,
		"organisation id":   p.OrganisationID,
		"type":              p.Type,
		"scheme":            p.Scheme,
		"scheme type":       p.SchemeType,
		"scheme sub-type":   p.SchemeSubType,
		"numeric reference": p.NumericRef,
	}
	for name, value := range fields {
		if value == "" {
			return errFieldBlank(name, value)
		}
	}
	return nil
}
