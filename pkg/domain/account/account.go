package account

import (
	"errors"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/bank"
)

type Account struct {
	Id      Identifier
	Name    string
	Address string
	Type    *int
	bankId  bank.Identifier
}

func (a Account) Validate() error {
	err := a.Id.Validate()
	if err != nil {
		return errNotValid(err, a.Id)
	}
	err = a.bankId.Validate()
	if err != nil {
		return errNotValid(err, a.Id)
	}
	if a.Name == "" {
		return errNotValid(errors.New("Name is blank"), a.Id)
	}
	if a.Address == "" {
		return errNotValid(errors.New("Name is blank"), a.Id)
	}
	return nil
}
