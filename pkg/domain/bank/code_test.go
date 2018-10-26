package bank

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCode_ToString(t *testing.T) {
	code := GBDSC
	assert.Equal(t, "GBDSC", code.String())
}

func TestCode_BBANFromString(t *testing.T) {
	code, err := ToCode("GBDSC")
	require.Nil(t, err)
	assert.Equal(t, GBDSC, *code)
}
func TestCode_Marshal(t *testing.T) {
	code := GBDSC
	b, err := json.Marshal(code)
	require.Nil(t, err)
	assert.Equal(t, `"GBDSC"`, string(b))
}

func TestCode_Unmarshal(t *testing.T) {
	var code Code
	err := json.Unmarshal([]byte(`"GBDSC"`), &code)
	require.Nil(t, err)
	assert.Equal(t, GBDSC, code)
}

func TestCode_UnmarshalError(t *testing.T) {
	var code Code
	err := json.Unmarshal([]byte(`"NOTVALID"`), &code)
	require.NotNil(t, err.(ErrCodeNotValid))
}

func TestCode_UnmarshalErrorNotString(t *testing.T) {
	var code Code
	err := json.Unmarshal([]byte("{"), &code)
	require.NotNil(t, err.(ErrCodeNotValid))
}
