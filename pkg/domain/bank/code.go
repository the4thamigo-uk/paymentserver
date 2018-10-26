package bank

import (
	"encoding/json"
)

// Code defines a bank code (https://www.ibm.com/support/knowledgecenter/en/SSRH46_3.0.0/fxhmapr2015clrcodes.html)
type Code int

const (
	// GBDSC is the UK domestic sort code
	GBDSC Code = iota
)

var (
	codesToStr = map[Code]string{
		GBDSC: "GBDSC",
	}
	codesFromStr = map[string]Code{}
)

func init() {
	for k, v := range codesToStr {
		codesFromStr[v] = k
	}
}

// String converts the bank code to a string
func (c *Code) String() string {
	return codesToStr[*c]
}

// ToCode attempts to read a code from a string
func ToCode(s string) (*Code, error) {
	c, ok := codesFromStr[s]
	if !ok {
		return nil, errCodeNotValid(s)
	}
	return &c, nil
}

// MarshalJSON implements the json.Marshaler interface for Code
func (c Code) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Code
func (c *Code) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errCodeNotValid(string(data))
	}
	c2, err := ToCode(s)
	if err != nil {
		return err
	}
	*c = *c2
	return nil
}
