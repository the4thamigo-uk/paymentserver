package money

import (
	"encoding/json"
	gomoney "github.com/Rhymond/go-money"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/amount"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/currency"
)

// Money defines an amount of a specified currency
type Money struct {
	m *gomoney.Money
}

type money struct {
	Amount   amount.Amount     `json:"amount"`
	Currency currency.Currency `json:"currency"`
}

// New creates money of the given amount in the given currency
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

// Equals checks if two money instances are equivalent
func (m Money) Equals(mny Money) bool {
	b, err := m.m.Equals(mny.m)
	return b && err == nil
}

// Currency returns the currency the money is denominated in
func (m Money) Currency() currency.Currency {
	return currency.Currency(m.m.Currency().Code)
}

// Places returns the number of decimal places
func (m Money) Places() int {
	return m.m.Currency().Fraction
}

// Amount returns the amount of the currency
func (m Money) Amount() amount.Amount {
	return amount.New(m.m.Amount(), m.m.Currency().Fraction)
}

// MarshalJSON implements the json.Marshaler interface for Money
func (m Money) MarshalJSON() ([]byte, error) {
	m2 := money{
		Amount:   m.Amount(),
		Currency: m.Currency(),
	}
	return json.Marshal(m2)
}

// UnmarshalJSON implements the json.Unmarshaler interface for Money
func (m *Money) UnmarshalJSON(data []byte) error {
	var m2 money
	err := json.Unmarshal(data, &m2)
	if err != nil {
		return err
	}
	m3, err := New(m2.Amount, m2.Currency.String())
	if err != nil {
		return err
	}
	*m = *m3
	return nil
}
