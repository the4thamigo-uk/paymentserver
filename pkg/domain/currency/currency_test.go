package currency

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCurrency_ToString(t *testing.T) {
	c := MustParse("GBP")
	assert.Equal(t, "GBP", c.String())
}

func TestCurrency_BBANFromString(t *testing.T) {
	c, err := Parse("GBP")
	require.Nil(t, err)
	assert.Equal(t, "GBP", c.String())
}

func TestCurrency_Marshal(t *testing.T) {
	c := MustParse("GBP")
	b, err := json.Marshal(c)
	require.Nil(t, err)
	assert.Equal(t, `"GBP"`, string(b))
}

func TestCurrency_Unmarshal(t *testing.T) {
	var c Currency
	err := json.Unmarshal([]byte(`"GBP"`), &c)
	require.Nil(t, err)
	assert.Equal(t, "GBP", c.String())
}

func TestCurrency_UnmarshalError(t *testing.T) {
	var c Currency
	err := json.Unmarshal([]byte(`"NOTVALID"`), &c)
	require.NotNil(t, err.(ErrCurrencyNotValid))
}

func TestCurrency_UnmarshalErrorNotString(t *testing.T) {
	var c Currency
	err := json.Unmarshal([]byte("{"), &c)
	require.NotNil(t, err.(ErrCurrencyNotValid))
}
