package account

import (
	"errors"
	"github.com/go-pascal/iban"
)

// Identifier defines a bank account number in a format compatible with the Code
type Identifier struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	Code   Code   `json:"code"`
	Type   *int   `json:"type,omitempty"`
}

// Validate returns whether the Identifier is valid.
// Currently only IBAN identifiers are actually validated.
func (id *Identifier) Validate() error {
	if id.Name == "" {
		return errIDNotValid(errors.New("Name is empty"), *id)
	}
	if id.Number == "" {
		return errIDNotValid(errors.New("Number is empty"), *id)
	}
	if id.Code == IBAN {
		_, err := iban.NewIBAN(id.Number)
		return errIDNotValid(err, *id)
	}
	return nil
}
