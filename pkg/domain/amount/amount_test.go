package amount

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAmount_ParsePositive(t *testing.T) {
	a, err := Parse("123.456")
	require.Nil(t, err)
	assert.Equal(t, 3, a.Precision())
	assert.Equal(t, int64(123456), a.Value())
	assert.Equal(t, "123.456", a.String())
}

func TestAmount_ParseNegative(t *testing.T) {
	a, err := Parse("-123.456")
	require.Nil(t, err)
	assert.Equal(t, 3, a.Precision())
	assert.Equal(t, int64(-123456), a.Value())
	assert.Equal(t, "-123.456", a.String())
}

func TestAmount_ParseMultipleDecimalPointsError(t *testing.T) {
	a, err := Parse("123.456.789")
	require.NotNil(t, err.(ErrAmountNotValid))
	assert.Equal(t, zero, a)
}

func TestAmount_ParseTrailingZeros(t *testing.T) {
	a, err := Parse("123.40")
	require.Nil(t, err)
	assert.Equal(t, 2, a.Precision())
	assert.Equal(t, "123.40", a.String())
}

func TestAmount_ParseZeroWithZeroPrecision(t *testing.T) {
	a, err := Parse("0")
	require.Nil(t, err)
	assert.Equal(t, "0", a.String())
}

func TestAmount_ParseZeroWithPrecision(t *testing.T) {
	a, err := Parse("0.00")
	require.Nil(t, err)
	assert.Equal(t, "0.00", a.String())
}

func TestAmount_ParseNonNumberError(t *testing.T) {
	a, err := Parse(" 123.456")
	require.NotNil(t, err.(ErrAmountNotValid))
	assert.Equal(t, zero, a)
}

func TestAmount_ParseMissingFractionError(t *testing.T) {
	a, err := Parse("123.")
	require.NotNil(t, err.(ErrAmountNotValid))
	assert.Equal(t, zero, a)
}

func TestAmount_ParseMissingIntegerError(t *testing.T) {
	a, err := Parse(".000")
	require.NotNil(t, err.(ErrAmountNotValid))
	assert.Equal(t, zero, a)
}

func TestAmount_ParseOnlyDecimalPointError(t *testing.T) {
	a, err := Parse(".")
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

func TestAmount_EmptyIsZero(t *testing.T) {
	var a Amount
	assert.Equal(t, "0", a.String())
}

func TestAmount_NewZeroPrecision(t *testing.T) {
	a := New(123456, 0)
	assert.Equal(t, "123456", a.String())
}

func TestAmount_NewNegativePrecision(t *testing.T) {
	a := New(123456, 0)
	assert.Equal(t, "123456", a.String())
}

func TestAmount_New(t *testing.T) {
	a := New(123456, 2)
	assert.Equal(t, "1234.56", a.String())
}

func TestAmount_RoundAwayFromZeroPositive(t *testing.T) {
	a := MustParse("1.567")
	assert.Equal(t, "2", a.Round(0).String())
	assert.Equal(t, "1.6", a.Round(1).String())
	assert.Equal(t, "1.57", a.Round(2).String())
}

func TestAmount_RoundAwayFromZeroNegative(t *testing.T) {
	a := MustParse("-1.567")
	assert.Equal(t, "-2", a.Round(0).String())
	assert.Equal(t, "-1.6", a.Round(1).String())
	assert.Equal(t, "-1.57", a.Round(2).String())
}

func TestAmount_RoundTowardsZeroPositive(t *testing.T) {
	a := MustParse("1.321")
	assert.Equal(t, "1", a.Round(0).String())
	assert.Equal(t, "1.3", a.Round(1).String())
	assert.Equal(t, "1.32", a.Round(2).String())
}

func TestAmount_RoundTowardsZeroNegative(t *testing.T) {
	a := MustParse("-1.321")
	assert.Equal(t, "-1", a.Round(0).String())
	assert.Equal(t, "-1.3", a.Round(1).String())
	assert.Equal(t, "-1.32", a.Round(2).String())
}

func TestAmount_RoundZero(t *testing.T) {
	a := MustParse("0.000")
	assert.Equal(t, "0.0", a.Round(1).String())
	assert.Equal(t, "0.00", a.Round(2).String())
}

func TestAmount_RoundToHigherPrecisionNoop(t *testing.T) {
	a := MustParse("123.456")
	assert.Equal(t, "123.456", a.Round(3).String())
	assert.Equal(t, "123.456", a.Round(4).String())
}

func TestAmount_RoundNegativePrecisionNoop(t *testing.T) {
	a := MustParse("123.456")
	assert.Equal(t, "123.456", a.Round(-1).String())
	assert.Equal(t, "123.456", a.Round(-2).String())
}

func TestAmount_MultiplyInteger(t *testing.T) {
	a := MustParse("123.455")
	b := MustParse("2")
	assert.Equal(t, "246.910", a.Multiply(b).String())
}

func TestAmount_MultiplyDifferentPrecision(t *testing.T) {
	a := MustParse("123.455")
	b := MustParse("0.555555")
	assert.Equal(t, "68.586042525", a.Multiply(b).String())
}

func TestAmount_MultiplyZero(t *testing.T) {
	a := MustParse("123.455")
	b := MustParse("0.0000")
	assert.Equal(t, "0.0000000", a.Multiply(b).String())
}

func TestAmount_MultiplyNegative(t *testing.T) {
	a := MustParse("123.455")
	b := MustParse("-1.234")
	assert.Equal(t, "-152.343470", a.Multiply(b).String())
}

func TestAmount_Marshal(t *testing.T) {
	a := New(12345, 2)
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
	err := json.Unmarshal([]byte("123"), &a)
	require.NotNil(t, err.(ErrAmountNotValid))
}
