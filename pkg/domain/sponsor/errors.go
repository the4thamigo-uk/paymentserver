package sponsor

import (
	"github.com/pkg/errors"
)

// ErrNumber indicates that the account number is incorrect
type ErrNumber error

func errNumber(num string) ErrNumber {
	return ErrNumber(errors.Errorf("The account number '%s' is not valid", num))
}

// ErrBankID indicates that the account is incorrect
type ErrBankID error

func errBankID(err error) ErrBankID {
	return ErrBankID(errors.Wrapf(err, "The bank details are not valid"))
}
