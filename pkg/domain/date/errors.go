package date

import (
	"github.com/pkg/errors"
)

// ErrDateParse indicates that the date is not valid
type ErrDateParse error

func errDateParse(err error, date string) ErrDateParse {
	return ErrDateParse(errors.Wrapf(err, "The date '%s' is not valid", date))
}
