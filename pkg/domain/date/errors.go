package date

import (
	"github.com/pkg/errors"
)

// ErrParseDate indicates that the date is not valid
type ErrParseDate error

func errParseDate(err error, date string) ErrParseDate {
	return ErrParseDate(errors.Wrapf(err, "The date '%s' is not valid", date))
}
