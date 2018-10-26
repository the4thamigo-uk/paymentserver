package charges

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCode_SHARToString(t *testing.T) {
	code := SHAR
	assert.Equal(t, "SHAR", code.String())
}

func TestCode_SHARFromString(t *testing.T) {
	code, err := Parse("SHAR")
	require.Nil(t, err)
	assert.Equal(t, SHAR, *code)
}

func TestCode_Marshal(t *testing.T) {
	code := SHAR
	b, err := json.Marshal(code)
	require.Nil(t, err)
	assert.Equal(t, `"SHAR"`, string(b))
}

func TestCode_Unmarshal(t *testing.T) {
	var code Code
	err := json.Unmarshal([]byte(`"SHAR"`), &code)
	require.Nil(t, err)
	assert.Equal(t, SHAR, code)
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
