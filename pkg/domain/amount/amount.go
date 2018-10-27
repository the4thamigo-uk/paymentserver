package amount

import (
	//"encoding/json"
	"github.com/shopspring/decimal"
)

// Amount is a fixed precision number
type Amount decimal.Decimal

var zero = Amount{}

// Parse extracts a currency code from the string
func Parse(amt string) (Amount, error) {
	r, err := decimal.NewFromString(amt)
	if err != nil {
		return zero, errAmountNotValid(err, amt)
	}
	return Amount(r), nil
}

// MustParse returns a valid Amount or otherwise panics
// Use for testing only
func MustParse(amt string) Amount {
	a, err := Parse(amt)
	if err != nil {
		panic(err)
	}
	return a
}

// New creates an amount to the power 10^exp
func New(amt int64, exp int) Amount {
	return Amount(decimal.New(amt, int32(exp)))
}

// String returns the Currency code as a string
func (a Amount) String() string {
	return decimal.Decimal(a).String()
}

func (a Amount) Multiply(b Amount) Amount {
	return Amount(decimal.Decimal(a).Mul(decimal.Decimal(b)))
}

// MarshalJSON implements the json.Marshaler interface for Amount
func (a Amount) MarshalJSON() ([]byte, error) {
	return decimal.Decimal(a).MarshalJSON()
}

// UnmarshalJSON implements the json.Unmarshaler interface for Currency
func (a *Amount) UnmarshalJSON(data []byte) error {
	err := (*decimal.Decimal)(a).UnmarshalJSON((data))
	if err != nil {
		return errAmountNotValid(err, string(data))
	}
	return nil
}
