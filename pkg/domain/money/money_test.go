package money

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMoney_MustParse(t *testing.T) {
	m := MustParse("123.45", "GBP")
	assert.Equal(t, "123.45", m.String())
	assert.Equal(t, 2, m.Places())
}

func TestMoney_MustParsePanics(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			require.NotNil(t, err.(ErrParseAmount))
		}
	}()
	_ = MustParse("", "GBP")
}

func TestMoney_InvalidMoneyError(t *testing.T) {
	m, err := Parse("1.0", "XXX")
	require.NotNil(t, err.(ErrCurrencyNotFound))
	assert.Nil(t, m)
}

func TestMoney_ParseEmptyError(t *testing.T) {
	m, err := Parse("", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseNonNumberError(t *testing.T) {
	m, err := Parse(" 123.45", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseBHD(t *testing.T) {
	m, err := Parse("123.456", "BHD")
	require.Nil(t, err)
	assert.Equal(t, "123.456", m.String())
	assert.Equal(t, 3, m.Places())
}

func TestMoney_ParseBHDNoUnit(t *testing.T) {
	m, err := Parse(".12", "BHD")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseBHDNoFractionError(t *testing.T) {
	m, err := Parse("123.", "BHD")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseBHDLongFractionLengthError(t *testing.T) {
	m, err := Parse("123.4567", "BHD")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseBHDShortFractionLengthError(t *testing.T) {
	m, err := Parse("123.45", "BHD")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseBHDMissingFractionError(t *testing.T) {
	m, err := Parse("123", "BHD")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseCLP(t *testing.T) {
	m, err := Parse("123", "CLP")
	require.Nil(t, err)
	assert.Equal(t, "123", m.String())
	assert.Equal(t, 0, m.Places())
}

func TestMoney_ParseCLPNoUnitError(t *testing.T) {
	m, err := Parse("", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseCLPHasFractionError(t *testing.T) {
	m, err := Parse("123.450", "CLP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseCLPMissingFractionError(t *testing.T) {
	m, err := Parse("123.", "CLP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_IsZeroError(t *testing.T) {
	m, err := Parse("0.00", "USD")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_IsNegativeError(t *testing.T) {
	m, err := Parse("-0.01", "USD")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_Money(t *testing.T) {
	m := MustParse("123.45", "GBP")
	assert.Equal(t, "GBP", m.Currency().String())
}

func TestMoney_Amount(t *testing.T) {
	m := MustParse("123.45", "GBP")
	assert.Equal(t, "123.45", m.Amount().String())
}

func TestMoney_AmountNegative(t *testing.T) {
	m := MustParse("123", "CLP")
	assert.Equal(t, "123", m.Amount().String())
}

func TestMoney_Equal(t *testing.T) {
	m1 := MustParse("123.45", "USD")
	m2 := MustParse("123.45", "USD")
	assert.True(t, m1.Equals(m2))
	assert.True(t, m2.Equals(m1))
}

func TestMoney_AmountNotEqual(t *testing.T) {
	m1 := MustParse("123.45", "USD")
	m2 := MustParse("123.46", "USD")
	assert.False(t, m1.Equals(m2))
	assert.False(t, m2.Equals(m1))
}

func TestMoney_MoneyNotEqual(t *testing.T) {
	m1 := MustParse("123.45", "USD")
	m2 := MustParse("123.45", "GBP")
	assert.False(t, m1.Equals(m2))
	assert.False(t, m2.Equals(m1))
}

func TestMoney_MarshalUnmarshalUSD(t *testing.T) {
	m1 := MustParse("123.45", "USD")
	b, err := json.Marshal(m1)
	require.Nil(t, err)
	var m2 Money
	err = json.Unmarshal(b, &m2)
	require.Nil(t, err)
	assert.True(t, m2.Equals(m1))
}

func TestMoney_MarshalUnmarshalCLP(t *testing.T) {
	m1 := MustParse("123", "CLP")
	b, err := json.Marshal(m1)
	require.Nil(t, err)
	var m2 Money
	err = json.Unmarshal(b, &m2)
	require.Nil(t, err)
	assert.True(t, m2.Equals(m1))
}

func TestMoney_UnmarshalError(t *testing.T) {
	var c Money
	err := json.Unmarshal([]byte(`{ "Amount": "123", "Currency": "USD"}`), &c)
	require.NotNil(t, err)
}

func TestMoney_UnmarshalErrorNotString(t *testing.T) {
	var c Money
	err := json.Unmarshal([]byte(`"NOTVALID"`), &c)
	require.NotNil(t, err)
}
