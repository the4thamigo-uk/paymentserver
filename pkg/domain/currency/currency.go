package currency

import (
	"encoding/json"
	gomoney "github.com/Rhymond/go-money"
)

// Currency defines an ISO4217 currency code
type Currency string

// Parse extracts a currency code from the string
func Parse(ccy string) (*Currency, error) {
	c := gomoney.GetCurrency(ccy)
	if c == nil {
		return nil, errCurrencyNotValid(ccy)
	}
	code := Currency(c.Code)
	return &code, nil
}

// MustParse returns a valid Currency or otherwise panics
// Use for testing only
func MustParse(ccy string) Currency {
	c, err := Parse(ccy)
	if err != nil {
		panic(err)
	}
	return *c
}

// String returns the Currency code as a string
func (c Currency) String() string {
	return string(c)
}

// MarshalJSON implements the json.Marshaler interface for Currency
func (c Currency) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Currency
func (c *Currency) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errCurrencyNotValid(string(data))
	}
	c2, err := Parse(s)
	if err != nil {
		return err
	}
	*c = *c2
	return nil
}
