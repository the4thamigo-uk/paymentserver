package account

import (
	"encoding/json"
)

type Code int

const (
	IBAN Code = iota
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

func (c *Code) String() string {
	return codesToStr[*c]
}

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
