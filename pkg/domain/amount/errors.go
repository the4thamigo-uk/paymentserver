package amount

import (
	"github.com/pkg/errors"
)

// ErrAmountNotValid indicates that the account code is incorrect
type ErrAmountNotValid error

func errAmountNotValid(err error, amt string) ErrAmountNotValid {
	return ErrAmountNotValid(errors.Wrapf(err, "The amount '%s' is not a valid", amt))
}
