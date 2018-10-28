package account

import (
	"github.com/pkg/errors"
)

// ErrCodeNotValid indicates that the account code is incorrect
type ErrCodeNotValid error

func errCodeNotValid(code string) ErrCodeNotValid {
	return ErrCodeNotValid(errors.Errorf("The account code '%s' is not a valid", code))
}

// ErrIDNotValid indicates that the account identifier is incorrect
type ErrIDNotValid error

func errIDNotValid(err error, id Identifier) ErrIDNotValid {
	return ErrIDNotValid(errors.Wrapf(err, "The account number '%s' is not a valid '%s'", id.Number, id.Code.String()))
}

// ErrNotValid indicates that the account is incorrect
type ErrNotValid error

func errNotValid(err error, id Identifier) ErrNotValid {
	return ErrNotValid(errors.Wrapf(err, "The account '%s' is not valid", id.Number))
}
