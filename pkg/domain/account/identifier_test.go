package account

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIdentifier_IBANBlankError(t *testing.T) {
	id := Identifier{"", IBAN}
	require.NotNil(t, id.Validate().(ErrIDNotValid))
}

func TestIdentifier_IBAN(t *testing.T) {
	id := Identifier{"GB82WEST12345698765432", IBAN}
	require.Nil(t, id.Validate())
}

func TestIdentifier_IBANBadChecksum(t *testing.T) {
	id := Identifier{"GB82WEST123lkjhl45698765432", IBAN}
	require.NotNil(t, id.Validate().(ErrIDNotValid))
}

func TestIdentifier_BBANBlankError(t *testing.T) {
	id := Identifier{"", BBAN}
	require.NotNil(t, id.Validate().(ErrIDNotValid))
}

func TestIdentifier_BBAN(t *testing.T) {
	id := Identifier{"NOTCHECKED", BBAN}
	require.Nil(t, id.Validate())
}
