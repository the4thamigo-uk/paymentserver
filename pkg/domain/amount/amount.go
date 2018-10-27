package amount

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Amount is a non-negative fixed-precision number
type Amount struct {
	val  int64
	prec int
}

var zero = Amount{}

// Parse extracts a currency code from the string
func Parse(amt string) (Amount, error) {
	parts := strings.Split(amt, ".")
	for _, s := range parts {
		if len(s) == 0 {
			return zero, errAmountNotValid(nil, amt)
		}
	}
	if len(parts) == 1 {
		parts = append(parts, "")
	}
	if len(parts) != 2 {
		return zero, errAmountNotValid(nil, amt)
	}
	prec := len(parts[1])
	s := strings.Join(parts, "")
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return zero, errAmountNotValid(nil, amt)
	}
	return New(val, prec), nil
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

// New creates an amount interpreting the last 'prec' digits of amt as decimals
func New(val int64, prec int) Amount {
	if prec < 0 {
		prec = 0
	}
	return Amount{
		val:  val,
		prec: prec,
	}
}

// Value returns the inner integer that holds the unscaled value
func (a Amount) Value() int64 {
	return a.val
}

// Precision returns the number of decimal places the Value should be scaled to
func (a Amount) Precision() int {
	return a.prec
}

// String returns the Currency code as a string
func (a Amount) String() string {
	if a.prec == 0 {
		return fmt.Sprintf("%d", a.val)
	}
	s := pow10(a.prec)
	return fmt.Sprintf("%d.%0"+fmt.Sprintf("%d", a.prec)+"d", a.val/s, abs(a.val)%s)
}

// Multiply multiples the Amount with b returning a new Amount
// with the same precision as a.
func (a Amount) Multiply(b Amount) Amount {
	// TODO check for int64 overflow
	return New(a.val*b.val, a.prec+b.prec)
}

// Round rounds the amount to the nearest prec decimal places
func (a Amount) Round(prec int) Amount {
	if prec >= a.prec || prec < 0 {
		return a
	}
	val := a.val / pow10(a.prec-prec-1)
	if a.val > 0 {
		val += 5
	} else if a.val < 0 {
		val -= 5
	}
	val /= 10
	return New(val, prec)
}

// MarshalJSON implements the json.Marshaler interface for Amount
func (a Amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Currency
func (a *Amount) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return errAmountNotValid(err, string(data))
	}
	a2, err := Parse(s)
	if err != nil {
		return errAmountNotValid(err, s)
	}
	*a = a2

	return nil
}

func pow10(exp int) int64 {
	s := int64(1)
	for i := 0; i < exp; i++ {
		s *= 10
	}
	return s
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}
