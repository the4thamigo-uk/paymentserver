package payment

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPayment_Validate(t *testing.T) {
	p := NewDummyPayment()
	err := p.Validate()
	require.Nil(t, err)
}

func TestPayment_BeneficiaryError(t *testing.T) {
	p := NewDummyPayment()
	p.Beneficiary.Name = ""
	err := p.Validate()
	require.NotNil(t, err.(ErrBeneficiaryNotValid))
}

func TestPayment_DebtorError(t *testing.T) {
	p := NewDummyPayment()
	p.Debtor.Name = ""
	err := p.Validate()
	require.NotNil(t, err.(ErrDebtorNotValid))
}

func TestPayment_SameAccountError(t *testing.T) {
	p := NewDummyPayment()
	p.Debtor.ID = p.Beneficiary.ID
	p.Debtor.BankID = p.Beneficiary.BankID
	err := p.Validate()
	t.Log(err)
	require.NotNil(t, err.(ErrSameAccount))
}

func TestPayment_ChargesError(t *testing.T) {
	p := NewDummyPayment()
	p.Charges.Sender = nil
	err := p.Validate()
	require.NotNil(t, err.(ErrChargesNotValid))
}

func TestPayment_NoFx(t *testing.T) {
	p := NewDummyPayment()
	p.Fx = nil
	err := p.Validate()
	require.Nil(t, err)
}

func TestPayment_FxError(t *testing.T) {
	p := NewDummyPayment()
	p.Fx.Reference = ""
	err := p.Validate()
	require.NotNil(t, err.(ErrChargesNotValid))
}

func TestPayment_SponsorError(t *testing.T) {
	p := NewDummyPayment()
	p.Sponsor.Number = ""
	err := p.Validate()
	require.NotNil(t, err.(ErrSponsorNotValid))
}

func TestPayment_IdBlankError(t *testing.T) {
	p := NewDummyPayment()
	p.ID = ""
	err := p.Validate()
	require.NotNil(t, err.(ErrFieldBlank))
}

func TestPayment_TypeBlankError(t *testing.T) {
	p := NewDummyPayment()
	p.Type = ""
	err := p.Validate()
	require.NotNil(t, err.(ErrFieldBlank))
}

func TestPayment_SchemeBlankError(t *testing.T) {
	p := NewDummyPayment()
	p.Scheme = ""
	err := p.Validate()
	require.NotNil(t, err.(ErrFieldBlank))
}

func TestPayment_SchemeTypeBlankError(t *testing.T) {
	p := NewDummyPayment()
	p.SchemeType = ""
	err := p.Validate()
	require.NotNil(t, err.(ErrFieldBlank))
}

func TestPayment_SchemeSubTypeBlankError(t *testing.T) {
	p := NewDummyPayment()
	p.SchemeSubType = ""
	err := p.Validate()
	require.NotNil(t, err.(ErrFieldBlank))
}

func TestPayment_NumericRefBlankError(t *testing.T) {
	p := NewDummyPayment()
	p.NumericRef = ""
	err := p.Validate()
	require.NotNil(t, err.(ErrFieldBlank))
}

func TestPayment_MarshalUnmarshalUSD(t *testing.T) {
	p1 := NewDummyPayment()
	b, err := json.Marshal(p1)
	require.Nil(t, err)
	var p2 Payment
	err = json.Unmarshal(b, &p2)
	require.Nil(t, err)
	assert.Equal(t, p2, p1)
}
