package money

import (
	gomoney "github.com/Rhymond/go-money"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/amount"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/currency"
)

// Money defines an amount of a specified currency
type Money struct {
	m *gomoney.Money
}

func New(amt amount.Amount, ccy string) (*Money, error) {
	c := gomoney.GetCurrency(ccy)
	if c == nil {
		return nil, errInvalidCurrency(ccy)
	}
	if amt.Precision() != c.Fraction {
		return nil, errParseAmount(amt.String(), ccy)
	}
	m := gomoney.New(amt.Value(), ccy)
	if !m.IsPositive() {
		return nil, errParseAmount(amt.String(), ccy)
	}
	return &Money{
		m: m,
	}, nil
}

// Parse returns a valid Money object for the given (positive) amount of currency.
// The amount is validated against the rules of the currency.
func Parse(amt string, ccy string) (*Money, error) {
	a, err := amount.Parse(amt)
	if err != nil {
		return nil, errParseAmount(amt, ccy)
	}
	return New(a, ccy)
}

// MustParse returns a valid Money or otherwise panics
// Use for testing only
func MustParse(amt string, ccy string) Money {
	m, err := Parse(amt, ccy)
	if err != nil {
		panic(err)
	}
	return *m
}

// String returns the string representation of the Money object
func (m Money) String() string {
	c := m.m.Currency()
	f := &gomoney.Formatter{
		Fraction: c.Fraction,
		Decimal:  c.Decimal,
		Template: "$1",
	}
	return f.Format(m.m.Amount())
}

// Currency returns the currency the money is denominated in
func (m Money) Currency() currency.Currency {
	return currency.Currency(m.m.Currency().Code)
}

// Amount returns the amount of the currency
func (m Money) Amount() amount.Amount {
	return amount.New(m.m.Amount(), m.m.Currency().Fraction)
}

func (m Money) Multiply(rate amount.Amount) (*Money, error) {
	a := m.Amount().Multiply(rate).Round(m.m.Currency().Fraction)
	return New(a, m.Currency().String())
}
