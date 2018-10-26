package money

import (
	"github.com/pkg/errors"
)

// ErrCurrencyNotFound indicates that the currency code is not recognised
type ErrCurrencyNotFound error

func errInvalidCurrency(ccy string) ErrCurrencyNotFound {
	return ErrCurrencyNotFound(errors.Errorf("Invalid currency '%s'", ccy))
}

// ErrParseAmount indicates that the amount is not valid for the currency
type ErrParseAmount error

func errParseAmount(amt string, ccy string) ErrParseAmount {
	return ErrParseAmount(errors.Errorf("The amount '%s' is not a valid amount for the currency '%s'", amt, ccy))
}
