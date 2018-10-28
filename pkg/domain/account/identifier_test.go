package account

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var accType = 0

func newTestIBAN() Identifier {
	// https://www.iban.com/structure
	return Identifier{
		Name:   "a name",
		Number: "IE64IRCE92050112345678",
		Code:   IBAN,
		Type:   &accType,
	}
}

func newTestBBAN() Identifier {
	return Identifier{
		Name:   "a name",
		Number: "NOTCHECKED",
		Code:   BBAN,
		Type:   &accType,
	}
}

func TestIdentifier_IBAN(t *testing.T) {
	id := newTestIBAN()
	require.Nil(t, id.Validate())
}

func TestIdentifier_IBANNameBlankError(t *testing.T) {
	id := newTestIBAN()
	id.Name = ""
	require.NotNil(t, id.Validate().(ErrIDNotValid))
}

func TestIdentifier_IBANNumberBlankError(t *testing.T) {
	id := newTestIBAN()
	id.Number = ""
	require.NotNil(t, id.Validate().(ErrIDNotValid))
}

func TestIdentifier_IBANBadChecksum(t *testing.T) {
	id := newTestIBAN()
	id.Number = "INVALID"
	require.NotNil(t, id.Validate().(ErrIDNotValid))
}

func TestIdentifier_BBAN(t *testing.T) {
	id := newTestBBAN()
	require.Nil(t, id.Validate())
}

func TestIdentifier_BBANNameBlankError(t *testing.T) {
	id := newTestBBAN()
	id.Name = ""
	require.NotNil(t, id.Validate().(ErrIDNotValid))
}

func TestIdentifier_BBANNumberBlankError(t *testing.T) {
	id := newTestBBAN()
	id.Number = ""
	require.NotNil(t, id.Validate().(ErrIDNotValid))
}

func TestIdentifier_ValidateNilType(t *testing.T) {
	a := newTestIBAN()
	a.Type = nil
	err := a.Validate()
	require.Nil(t, err)
}
