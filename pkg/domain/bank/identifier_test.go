package bank

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIdentifier_IdError(t *testing.T) {
	id := Identifier{"", GBDSC}
	require.NotNil(t, id.Validate().(ErrBankIDNotValid))
}

func TestIdentifier_Id(t *testing.T) {
	id := Identifier{"i123", GBDSC}
	require.Nil(t, id.Validate())
}
