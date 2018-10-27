package money

import (
	gomoney "github.com/Rhymond/go-money"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/currency"
	"strconv"
	"strings"
)

// Money defines an amount of a specified currency
type Money struct {
	m *gomoney.Money
}

var (
	zero = Money{}
)

// Parse returns a valid Money object for the given (positive) amount of currency.
// The amount is validated against the rules of the currency.
func Parse(amt string, ccy string) (Money, error) {
	c := gomoney.GetCurrency(ccy)
	if c == nil {
		return zero, errInvalidCurrency(ccy)
	}
	parts := strings.Split(amt, ".")
	if c.Fraction == 0 {
		if len(parts) != 1 {
			return zero, errParseAmount(amt, ccy)
		}
		parts = append(parts, "")
	}
	if len(parts) != 2 {
		return zero, errParseAmount(amt, ccy)
	}
	if len(parts[1]) != c.Fraction {
		return zero, errParseAmount(amt, ccy)
	}
	if len(parts[0]) == 0 {
		return zero, errParseAmount(amt, ccy)
	}
	s := strings.Join(parts, "")
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return zero, errParseAmount(amt, ccy)
	}
	m := gomoney.New(n, ccy)
	if !m.IsPositive() {
		return zero, errParseAmount(amt, ccy)
	}
	return Money{
		m: m,
	}, nil
}

// MustParse returns a valid Money or otherwise panics
// Use for testing only
func MustParse(amt string, ccy string) Money {
	m, err := Parse(amt, ccy)
	if err != nil {
		panic(err)
	}
	return m
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
