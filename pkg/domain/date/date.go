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
		return zero, errDateParse(err, s)
	}
	return Date(t), nil
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
		return errDateParse(err, string(data))
	}
	d2, err := Parse(s)
	if err != nil {
		return err
	}
	*d = d2
	return nil
}
