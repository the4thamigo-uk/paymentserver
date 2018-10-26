package account

import (
	"errors"
	"github.com/go-pascal/iban"
)

type Identifier struct {
	Number string
	Code   Code
}

// Validate returns whether the Identifier is valid.
// Currently only IBAN identifiers are actually validated.
func (id Identifier) Validate() error {
	if id.Code == IBAN {
		_, err := iban.NewIBAN(id.Number)
		return errIdNotValid(err, id)
	}
	if id.Number == "" {
		return errIdNotValid(errors.New("Number is empty"), id)
	}
	return nil
}
