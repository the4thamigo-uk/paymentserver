package payment

import (
	"github.com/pkg/errors"
)

// ErrBeneficiaryNotValid indicates that the beneficiary is incorrect
type ErrBeneficiaryNotValid error

func errBeneficiaryNotValid(err error) ErrBeneficiaryNotValid {
	return ErrBeneficiaryNotValid(errors.Wrapf(err, "The beneficiary is not valid"))
}

// ErrDebtorNotValid indicates that the debtor is incorrect
type ErrDebtorNotValid error

func errDebtorNotValid(err error) ErrDebtorNotValid {
	return ErrDebtorNotValid(errors.Wrapf(err, "The debtor is not valid"))
}

// ErrSameAccount indicates that the debtor is incorrect
type ErrSameAccount error

func errSameAccount() ErrSameAccount {
	return ErrSameAccount(errors.New("The beneficiary has the same account as the debtor"))
}

// ErrChargesNotValid indicates that the charges are incorrect
type ErrChargesNotValid error

func errChargesNotValid(err error) ErrChargesNotValid {
	return ErrChargesNotValid(errors.Wrapf(err, "The charges are not valid"))
}

// ErrFxNotValid indicates that the foreign exchange contract is incorrect
type ErrFxNotValid error

func errFxNotValid(err error) ErrFxNotValid {
	return ErrFxNotValid(errors.Wrapf(err, "The foreign exchange contract is not valid"))
}

// ErrSponsorNotValid indicates that the sponsor is incorrect
type ErrSponsorNotValid error

func errSponsorNotValid(err error) ErrSponsorNotValid {
	return ErrSponsorNotValid(errors.Wrapf(err, "The sponsor is not valid"))
}

// ErrFieldBlank indicates that a mandatory payment field value is blank
type ErrFieldBlank error

func errFieldBlank(name string, value interface{}) ErrFieldBlank {
	return ErrFieldBlank(errors.Errorf("The payment attribute '%s' with value '%v' is not valid", name, value))
}
