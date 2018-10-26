package date

import (
	"encoding/json"
	"time"
)

// Date represents a pure date in YYYY-MM-DD format.
type Date time.Time

const layout = "2006-01-02"

var zero = Date(time.Time{})

// Parse extracts a date from a string
func Parse(s string) (Date, error) {
	t, err := time.Parse(layout, s)
	if err != nil {
		return zero, errParseDate(err, s)
	}
	return Date(t), nil
}

// MustParse returns a valid Date or otherwise panics
// Use for testing only
func MustParse(s string) Date {
	d, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return d
}

// String returns the data in the standard format
func (d *Date) String() string {
	return time.Time(*d).Format(layout)
}

// MarshalJSON implements the json.Marshaler interface for Date
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Date
func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errParseDate(err, string(data))
	}
	d2, err := Parse(s)
	if err != nil {
		return err
	}
	*d = d2
	return nil
}
