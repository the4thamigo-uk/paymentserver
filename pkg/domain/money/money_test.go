package money

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMoney_MustParse(t *testing.T) {
	m := MustParse("123.45", "GBP")
	assert.Equal(t, "123.45", m.String())
}

func TestMoney_MustParsePanics(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			require.NotNil(t, err.(ErrParseAmount))
		}
	}()
	_ = MustParse("", "GBP")
}

func TestMoney_InvalidCurrencyError(t *testing.T) {
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

func TestMoney_ParseGBP(t *testing.T) {
	m, err := Parse("123.45", "GBP")
	require.Nil(t, err)
	assert.Equal(t, "123.45", m.String())
}

func TestMoney_ParseGBPNoUnit(t *testing.T) {
	m, err := Parse(".12", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseGBPNoFractionError(t *testing.T) {
	m, err := Parse("123.", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseGBPLongFractionLengthError(t *testing.T) {
	m, err := Parse("123.456", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseGBPShortFractionLengthError(t *testing.T) {
	m, err := Parse("123.4", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseGBPMissingFractionError(t *testing.T) {
	m, err := Parse("123", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Nil(t, m)
}

func TestMoney_ParseCLP(t *testing.T) {
	m, err := Parse("123", "CLP")
	require.Nil(t, err)
	assert.Equal(t, "123", m.String())
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

func TestMoney_Currency(t *testing.T) {
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

func TestMoney_CurrencyNotEqual(t *testing.T) {
	m1 := MustParse("123.45", "USD")
	m2 := MustParse("123.45", "GBP")
	assert.False(t, m1.Equals(m2))
	assert.False(t, m2.Equals(m1))
}
