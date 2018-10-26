package account

import (
	"errors"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/bank"
)

// Account defines a bank account
type Account struct {
	ID      Identifier
	Name    string
	Address string
	Type    *int
	bankID  bank.Identifier
}

// Validate performs some basic checks on the validity of the Account
func (a Account) Validate() error {
	err := a.ID.Validate()
	if err != nil {
		return errNotValid(err, a.ID)
	}
	err = a.bankID.Validate()
	if err != nil {
		return errNotValid(err, a.ID)
	}
	if a.Name == "" {
		return errNotValid(errors.New("Name is blank"), a.ID)
	}
	if a.Address == "" {
		return errNotValid(errors.New("Name is blank"), a.ID)
	}
	return nil
}
