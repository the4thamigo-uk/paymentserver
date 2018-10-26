package account

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCode_BBANToString(t *testing.T) {
	code := BBAN
	assert.Equal(t, "BBAN", code.String())
}

func TestCode_BBANFromString(t *testing.T) {
	code, err := ToCode("BBAN")
	require.Nil(t, err)
	assert.Equal(t, BBAN, *code)
}

func TestCode_Marshal(t *testing.T) {
	code := BBAN
	b, err := json.Marshal(code)
	require.Nil(t, err)
	assert.Equal(t, `"BBAN"`, string(b))
}

func TestCode_Unmarshal(t *testing.T) {
	var code Code
	err := json.Unmarshal([]byte(`"BBAN"`), &code)
	require.Nil(t, err)
	assert.Equal(t, BBAN, code)
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
