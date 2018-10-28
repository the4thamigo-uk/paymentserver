package amount

import (
	"github.com/pkg/errors"
)

// ErrAmountNotValid indicates that the amount is not valid
type ErrAmountNotValid error

func errAmountNotValid(err error, amt string) ErrAmountNotValid {
	if err != nil {
		return ErrAmountNotValid(errors.Wrapf(err, "The amount '%s' is not valid", amt))
	}
	return ErrAmountNotValid(errors.Errorf("The amount '%s' is not valid", amt))
}
