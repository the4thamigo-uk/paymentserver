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
	m, err := Parse("", "XXX")
	require.NotNil(t, err.(ErrCurrencyNotFound))
	assert.Equal(t, zero, m)
}

func TestMoney_ParseEmptyError(t *testing.T) {
	m, err := Parse("", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Equal(t, zero, m)
}

func TestMoney_ParseNonNumberError(t *testing.T) {
	m, err := Parse(" 123.45", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Equal(t, zero, m)
}

func TestMoney_ParseGBP(t *testing.T) {
	m, err := Parse("123.45", "GBP")
	require.Nil(t, err)
	assert.Equal(t, "123.45", m.String())
}

func TestMoney_ParseGBPNoUnitError(t *testing.T) {
	m, err := Parse(".12", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Equal(t, zero, m)
}

func TestMoney_ParseGBPNoFractionError(t *testing.T) {
	m, err := Parse("123.", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Equal(t, zero, m)
}

func TestMoney_ParseGBPLongFractionLengthError(t *testing.T) {
	m, err := Parse("123.456", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Equal(t, zero, m)
}

func TestMoney_ParseGBPShortFractionLengthError(t *testing.T) {
	m, err := Parse("123.4", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Equal(t, zero, m)
}

func TestMoney_ParseGBPMissingFractionError(t *testing.T) {
	m, err := Parse("123", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Equal(t, zero, m)
}

func TestMoney_ParseCLP(t *testing.T) {
	m, err := Parse("123", "CLP")
	require.Nil(t, err)
	assert.Equal(t, "123", m.String())
}

func TestMoney_ParseCLPNoUnitError(t *testing.T) {
	m, err := Parse("", "GBP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Equal(t, zero, m)
}

func TestMoney_ParseCLPHasFractionError(t *testing.T) {
	m, err := Parse("123.450", "CLP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Equal(t, zero, m)
}

func TestMoney_ParseCLPMissingFractionError(t *testing.T) {
	m, err := Parse("123.", "CLP")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Equal(t, zero, m)
}

func TestMoney_IsNotPositiveError(t *testing.T) {
	m, err := Parse("0.00", "USD")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Equal(t, zero, m)
}

func TestMoney_IsNegativeError(t *testing.T) {
	m, err := Parse("-0.01", "USD")
	require.NotNil(t, err.(ErrParseAmount))
	assert.Equal(t, zero, m)
}
