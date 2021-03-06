package account

import (
	"encoding/json"
)

// Code defines a bank account code type type Code int
type Code int

const (
	// IBAN is an International Bank Account Number
	IBAN Code = iota
	// BBAN is a Basic Bank Account Number
	BBAN
)

var (
	codesToStr = map[Code]string{
		IBAN: "IBAN",
		BBAN: "BBAN",
	}
	codesFromStr = map[string]Code{}
)

func init() {
	for k, v := range codesToStr {
		codesFromStr[v] = k
	}
}

// String returns a string representation of the Code
func (c *Code) String() string {
	return codesToStr[*c]
}

// Parse attempts to parse a Code from a string
//TODO rename to ParseCode
func Parse(s string) (*Code, error) {
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
	c2, err := Parse(s)
	if err != nil {
		return err
	}
	*c = *c2
	return nil
}
