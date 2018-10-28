package fx

import (
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/amount"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
)

// Contract specifies the details of the foreign exchange contract.
type Contract struct {
	Reference string        `json:"ref"`
	Rate      amount.Amount `json:"rate"` // the exchange rate to convert foreign to domestic
	Domestic  money.Money   `json:"domestic"`
}

// Validate performs some checks on the Contract.
func (c *Contract) Validate(foreign money.Money) error {
	if c.Reference == "" {
		return errReferenceNotValid()
	}
	domestic, err := c.Exchange(foreign)
	if err != nil {
		return errRateNotValid(c.Rate, foreign, c.Domestic)
	}
	if !domestic.Equals(c.Domestic) {
		return errRateNotValid(c.Rate, foreign, c.Domestic)
	}
	return nil
}

// Exchange computes the domestic money from the given foreign money at the contract rate
func (c *Contract) Exchange(foreign money.Money) (*money.Money, error) {
	forAmt := foreign.Amount()
	domAmt := forAmt.Multiply(c.Rate).Round(c.Domestic.Places())
	dom, err := money.New(domAmt, c.Domestic.Currency().String())
	if err != nil {
		return nil, errExchange(c.Rate, foreign)
	}
	return dom, nil
}
