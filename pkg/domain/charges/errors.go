package charges

import (
	"github.com/pkg/errors"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/currency"
)

// ErrCodeNotValid indicates that the bearer code is incorrect
type ErrCodeNotValid error

func errCodeNotValid(code string) ErrCodeNotValid {
	return ErrCodeNotValid(errors.Errorf("The bearer code '%s' is not valid", code))
}

// ErrChargeNotFound indicates that a sender or receiver is not found for the given currency
type ErrChargeNotFound error

func errChargeNotFound(side string, ccy currency.Currency) ErrChargeNotFound {
	return ErrChargeNotFound(errors.Errorf("There is no '%s' charge for currency '%s'", side, ccy))
}

// ErrNumberCharges indicates that the number of sender charges is incorrect
type ErrNumberCharges error

func errNumberCharges(n int) ErrNumberCharges {
	return ErrNumberCharges(errors.Errorf("The number of sender charges (%d) is incorrect", n))
}
