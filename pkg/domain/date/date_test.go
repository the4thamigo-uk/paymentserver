package date

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDate_Parse(t *testing.T) {
	d, err := Parse("2000-01-02")
	require.Nil(t, err)
	assert.Equal(t, "2000-01-02", d.String())
}

func TestDate_Marshal(t *testing.T) {
	d, err := Parse("2000-01-02")
	require.Nil(t, err)
	b, err := json.Marshal(d)
	require.Nil(t, err)
	assert.Equal(t, `"2000-01-02"`, string(b))
}

func TestDate_Unmarshal(t *testing.T) {
	var d Date
	err := json.Unmarshal([]byte(`"2000-01-02"`), &d)
	require.Nil(t, err)
	assert.Equal(t, "2000-01-02", d.String())
}

func TestDate_UnmarshalError(t *testing.T) {
	var d Date
	err := json.Unmarshal([]byte(`"NOTVALID"`), &d)
	require.NotNil(t, err.(ErrDateParse))
}

func TestDate_UnmarshalErrorNotString(t *testing.T) {
	var d Date
	err := json.Unmarshal([]byte("{"), &d)
	require.NotNil(t, err.(ErrDateParse))
}
