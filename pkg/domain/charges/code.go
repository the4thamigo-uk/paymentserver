package charges

import (
	"encoding/json"
)

// Code defines a bank account code type type Code int
type Code int

const (
	// SHAR indicates charges are shared between the debtor and the beneficiary
	SHAR Code = iota
	// CRED indicates charges are borne by the beneficiary/creditor
	CRED
	// DEBT indicates charges are borne by the debtor
	DEBT
	// SLEV special code for SEPA credit transfers (not sure this is relevant
	//SLEV
)

var (
	codesToStr = map[Code]string{
		SHAR: "SHAR",
		CRED: "CRED",
		DEBT: "DEBT",
		//SLEV: "SLEV",
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
