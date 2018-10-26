package charges

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCharges_Validate(t *testing.T) {
	a := newTest()
	err := a.Validate()
	require.Nil(t, err)
}
