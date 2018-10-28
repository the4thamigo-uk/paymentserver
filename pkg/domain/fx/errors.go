package fx

import (
	"github.com/pkg/errors"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/amount"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
)

// ErrReferenceNotValid indicates that the contract reference is incorrect
type ErrReferenceNotValid error

func errReferenceNotValid() ErrReferenceNotValid {
	return ErrReferenceNotValid(errors.Errorf("The contract reference is not valid"))
}

// ErrExchange indicates that an exchange could not be calculated for the given money
type ErrExchange error

func errExchange(rate amount.Amount, foreign money.Money) ErrExchange {
	return ErrExchange(errors.Errorf("Failed to calculate a valid exchange from %s (%s) at rate %s",
		foreign.String(),
		foreign.Currency(),
		rate.String()))
}

// ErrRateNotValid indicates that the computed exchange from foreign money to domestic money is not correct
type ErrRateNotValid error

func errRateNotValid(rate amount.Amount, foreign money.Money, domestic money.Money) ErrRateNotValid {
	return ErrRateNotValid(errors.Errorf("The exchange rate '%s' from %s (%s) to %s (%s) is not correct",
		rate.String(),
		foreign.String(),
		foreign.Currency(),
		domestic.String(),
		domestic.Currency()))
}
