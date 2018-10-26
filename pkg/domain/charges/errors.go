package charges

import (
	"github.com/pkg/errors"
)

// ErrCodeNotValid indicates that the account code is incorrect
type ErrCodeNotValid error

func errCodeNotValid(code string) ErrCodeNotValid {
	return ErrCodeNotValid(errors.Errorf("The bearer code '%s' is not valid", code))
}
