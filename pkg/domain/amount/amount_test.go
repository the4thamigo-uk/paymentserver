package amount

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAmount_Parse(t *testing.T) {
	a, err := Parse("123.456")
	require.Nil(t, err)
	assert.Equal(t, "123.456", a.String())
}

func TestAmount_ParseNegative(t *testing.T) {
	a, err := Parse("-123.456")
	require.Nil(t, err)
	assert.Equal(t, "-123.456", a.String())
}

func TestAmount_ParsecwZeroNegative(t *testing.T) {
	a, err := Parse("0")
	require.Nil(t, err)
	assert.Equal(t, "0", a.String())
}

func TestAmount_ParseError(t *testing.T) {
	a, err := Parse(" 123.456")
	require.NotNil(t, err.(ErrAmountNotValid))
	assert.Equal(t, zero, a)
}

func TestAmount_MustParsePanics(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			require.NotNil(t, err.(ErrAmountNotValid))
		}
	}()
	_ = MustParse("NOTVALID")
}

func TestAmount_NewZeroExp(t *testing.T) {
	a := New(123456, 0)
	assert.Equal(t, "123456", a.String())
}

func TestAmount_NewPositiveExp(t *testing.T) {
	a := New(123456, 2)
	assert.Equal(t, "12345600", a.String())
}

func TestAmount_NewNegativeExponent(t *testing.T) {
	a := New(123456, -3)
	assert.Equal(t, "123.456", a.String())
}

func TestAmount_Multiply(t *testing.T) {
	a := MustParse("123.455")
	b := MustParse("0.5")
	assert.Equal(t, "61.7275", a.Multiply(b).String())
}

func TestAmount_Marshal(t *testing.T) {
	a := New(12345, -2)
	b, err := json.Marshal(a)
	require.Nil(t, err)
	assert.Equal(t, `"123.45"`, string(b))
}

func TestAmount_Unmarshal(t *testing.T) {
	var a Amount
	err := json.Unmarshal([]byte(`"123.45"`), &a)
	require.Nil(t, err)
	assert.Equal(t, "123.45", a.String())
}

func TestAmount_UnmarshalError(t *testing.T) {
	var a Amount
	err := json.Unmarshal([]byte(`"NOTVALID"`), &a)
	require.NotNil(t, err.(ErrAmountNotValid))
}

func TestAmount_UnmarshalErrorNotString(t *testing.T) {
	var a Amount
	err := json.Unmarshal([]byte("{"), &a)
	require.NotNil(t, err.(ErrAmountNotValid))
}
