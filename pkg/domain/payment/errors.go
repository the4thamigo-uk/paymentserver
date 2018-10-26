package payment

import (
	"github.com/pkg/errors"
)

// ErrBeneficiaryNotValid indicates that the debtor is incorrect
type ErrBeneficiaryNotValid error

func errBeneficiaryNotValid(err error) ErrBeneficiaryNotValid {
	return ErrBeneficiaryNotValid(errors.Wrapf(err, "The debtor is not valid"))
}

// ErrDebtorNotValid indicates that the debtor is incorrect
type ErrDebtorNotValid error

func errDebtorNotValid(err error) ErrDebtorNotValid {
	return ErrDebtorNotValid(errors.Wrapf(err, "The debtor is not valid"))
}
