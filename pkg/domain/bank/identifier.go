package bank

import (
	"errors"
)

// Identifier is a bank id in a format defined by the bank code
// TODO: need to implement https://wiki.xmldation.com/Support/ISO20022/General_Rules/Clearing_codes
type Identifier struct {
	ID   string `json:"id"`
	Code Code   `json:"code"`
}

// Validate returns whether the Identifier is valid.
// Currently only performs basic checks.
func (id Identifier) Validate() error {
	if id.ID == "" {
		return errBankIDNotValid(errors.New("Bank id is empty"), id)
	}
	return nil
}
