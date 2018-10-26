package account

import (
	"github.com/pkg/errors"
)

// ErrCodeNotValid indicates that the account code is incorrect
type ErrCodeNotValid error

func errCodeNotValid(code string) ErrCodeNotValid {
	return ErrCodeNotValid(errors.Errorf("The account code '%s' is not a valid", code))
}

// ErrIdNotValid indicates that the account identifier is incorrect
type ErrIdNotValid error

func errIdNotValid(err error, id Identifier) ErrIdNotValid {
	return ErrIdNotValid(errors.Wrapf(err, "The account number '%s' is not a valid '%s'", id.Number, id.Code))
}

// ErrNotValid indicates that the account is incorrect
type ErrNotValid error

func errNotValid(err error, id Identifier) ErrNotValid {
	return ErrNotValid(errors.Wrapf(err, "The account '%s' is not valid '%s'", id.Number))
}
