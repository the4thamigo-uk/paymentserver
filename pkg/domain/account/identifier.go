package account

import (
	"errors"
	"github.com/go-pascal/iban"
)

// Identifier defines a bank account number in a format compatible with the Code
type Identifier struct {
	Number string `json:"number"`
	Code   Code   `json:"code"`
}

// Validate returns whether the Identifier is valid.
// Currently only IBAN identifiers are actually validated.
func (id Identifier) Validate() error {
	if id.Code == IBAN {
		_, err := iban.NewIBAN(id.Number)
		return errIDNotValid(err, id)
	}
	if id.Number == "" {
		return errIDNotValid(errors.New("Number is empty"), id)
	}
	return nil
}
