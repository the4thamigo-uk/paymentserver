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

// ErrChargesNotValid indicates that the debtor is incorrect
type ErrChargesNotValid error

func errChargesNotValid(err error) ErrChargesNotValid {
	return ErrChargesNotValid(errors.Wrapf(err, "The charges are not valid"))
}

// ErrFxNotValid indicates that the debtor is incorrect
type ErrFxNotValid error

func errFxNotValid(err error) ErrFxNotValid {
	return ErrFxNotValid(errors.Wrapf(err, "The foreign exchange contract is not valid"))
}
