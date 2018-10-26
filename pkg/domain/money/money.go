package money

import (
	gomoney "github.com/Rhymond/go-money"
	"strconv"
	"strings"
)

type Money struct {
	m *gomoney.Money
}

var (
	zero = Money{}
)

// Parse returns a valid Money object for the given amount and currency.
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
	c = m.Currency()
	return Money{
		m: gomoney.New(n, ccy),
	}, nil
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