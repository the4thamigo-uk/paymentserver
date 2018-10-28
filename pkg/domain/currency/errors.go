package currency

import (
	"github.com/pkg/errors"
)

// ErrCurrencyNotValid indicates that the currency code is incorrect
type ErrCurrencyNotValid error

func errCurrencyNotValid(code string) ErrCurrencyNotValid {
	return ErrCurrencyNotValid(errors.Errorf("The currency code '%s' is not a valid", code))
}
