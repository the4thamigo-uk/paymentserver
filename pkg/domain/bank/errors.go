package bank

import (
	"github.com/pkg/errors"
)

// ErrCodeNotValid indicates that the account code is incorrect
type ErrCodeNotValid error

func errCodeNotValid(code string) ErrCodeNotValid {
	return ErrCodeNotValid(errors.Errorf("The account code '%s' is not a valid", code))
}

// ErrBankIDNotValid indicates that the bank identifier is incorrect
type ErrBankIDNotValid error

func errBankIDNotValid(err error, id Identifier) ErrBankIDNotValid {
	return ErrBankIDNotValid(errors.Wrapf(err, "The bank id '%s' is not a valid '%v'", id.ID, id.Code))
}