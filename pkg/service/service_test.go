package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/payment"
	"github.com/the4thamigo-uk/paymentserver/pkg/presentation"
	"github.com/the4thamigo-uk/paymentserver/pkg/store/memorystore"
	"testing"
)

func TestService_Create(t *testing.T) {
	s := memorystore.New()
	dp1 := payment.NewDummyPayment()
	pp1, err := presentation.FromDomainPayment(dp1)
	require.Nil(t, err)
	pp2, err := CreatePayment(s, pp1)
	require.Nil(t, err)
	assert.NotEqual(t, dp1.Entity.ID, pp2.ID)
	assert.Equal(t, 1, pp2.Version)
}

func TestService_InitialSaveVersion0(t *testing.T) {
	s := memorystore.New()
	dp1 := payment.NewDummyPayment()
	dp1.Entity.Version = 0
	pp1, err := presentation.FromDomainPayment(dp1)
	require.Nil(t, err)
	pp2, err := SavePayment(s, pp1)
	require.Nil(t, err)
	assert.Equal(t, 1, pp2.Version)
	pp1.Version = 1
	assert.Equal(t, *pp1, *pp2)
}

func TestService_InitialSaveVersionNonZero(t *testing.T) {
	s := memorystore.New()
	dp1 := payment.NewDummyPayment()
	dp1.Entity.Version = 1
	pp1, err := presentation.FromDomainPayment(dp1)
	require.Nil(t, err)
	pp2, err := SavePayment(s, pp1)
	require.Nil(t, err)
	assert.Equal(t, 2, pp2.Version)
	pp1.Version = 2
	assert.Equal(t, *pp1, *pp2)
}

func TestService_SaveValidationError(t *testing.T) {
	s := memorystore.New()
	dp1 := payment.NewDummyPayment()
	pp1, err := presentation.FromDomainPayment(dp1)
	require.Nil(t, err)

	// try to save an invalid payment
	pp1.Attributes.BeneficiaryParty = presentation.Account{}
	_, err = SavePayment(s, pp1)
	require.NotNil(t, err)
}

func TestService_Load(t *testing.T) {
	s := memorystore.New()
	dp1 := payment.NewDummyPayment()
	pp1, err := presentation.FromDomainPayment(dp1)
	require.Nil(t, err)
	pp2, err := SavePayment(s, pp1)
	require.Nil(t, err)

	eid := presentation.Entity{
		ID: dp1.Entity.ID,
	}
	pp3, err := LoadPayment(s, eid)
	require.Nil(t, err)
	assert.Equal(t, 1, pp3.Version)
	assert.Equal(t, *pp3, *pp2)
}

func TestService_LoadWrongID(t *testing.T) {
	s := memorystore.New()
	eid := presentation.Entity{
		ID: "WRONGID",
	}
	_, err := LoadPayment(s, eid)
	require.NotNil(t, err)
}

func TestService_Delete(t *testing.T) {
	s := memorystore.New()
	dp1 := payment.NewDummyPayment()
	pp1, err := presentation.FromDomainPayment(dp1)
	require.Nil(t, err)
	pp2, err := SavePayment(s, pp1)
	require.Nil(t, err)

	eid := presentation.Entity{
		ID: dp1.Entity.ID,
	}
	pp3, err := DeletePayment(s, eid)
	require.Nil(t, err)
	assert.Equal(t, 1, pp3.Version)
	assert.Equal(t, *pp3, *pp2)

	_, err = LoadPayment(s, eid)
	require.NotNil(t, err)
}

func TestService_DeleteWrongID(t *testing.T) {
	s := memorystore.New()
	eid := presentation.Entity{
		ID: "WRONGID",
	}
	_, err := DeletePayment(s, eid)
	require.NotNil(t, err)

	_, err = LoadPayment(s, eid)
	require.NotNil(t, err)
}

func TestService_SaveLatest(t *testing.T) {
	s := memorystore.New()
	dp1 := payment.NewDummyPayment()
	pp1, err := presentation.FromDomainPayment(dp1)
	require.Nil(t, err)
	_, err = SavePayment(s, pp1)
	require.Nil(t, err)

	// Overwrite latest version (1) using version 0
	pp1.Version = 0
	pp1.Attributes.Reference = "MODIFIED"
	pp3, err := SavePayment(s, pp1)
	assert.Equal(t, "MODIFIED", pp3.Attributes.Reference)
	require.Nil(t, err)
	assert.Equal(t, 2, pp3.Version)
	pp1.Version = 2
	assert.Equal(t, *pp3, *pp1)
}

func TestService_SaveWrongVersion(t *testing.T) {
	s := memorystore.New()
	dp1 := payment.NewDummyPayment()
	pp1, err := presentation.FromDomainPayment(dp1)
	require.Nil(t, err)
	_, err = SavePayment(s, pp1)
	require.Nil(t, err)

	// Save wrong version yields error
	pp1.Version = 100
	_, err = SavePayment(s, pp1)
	require.NotNil(t, err)
}

func TestService_Update(t *testing.T) {
	s := memorystore.New()
	dp1 := payment.NewDummyPayment()
	pp1, err := presentation.FromDomainPayment(dp1)
	require.Nil(t, err)
	_, err = SavePayment(s, pp1)
	require.Nil(t, err)

	// create a patch
	pp2 := *pp1
	pp2.Attributes = presentation.Attributes{}
	pp2.Attributes.BeneficiaryParty.AccountName = "DIFFERENT"
	pp3, err := UpdatePayment(s, &pp2)
	require.Nil(t, err)
	pp1.Attributes.BeneficiaryParty.AccountName = "DIFFERENT"
	pp1.Version = 2
	assert.Equal(t, *pp1, *pp3)
}

func TestService_UpdateWrongID(t *testing.T) {
	s := memorystore.New()
	dp1 := payment.NewDummyPayment()
	pp1, err := presentation.FromDomainPayment(dp1)
	require.Nil(t, err)

	_, err = UpdatePayment(s, pp1)
	require.NotNil(t, err)
}

func TestService_List(t *testing.T) {
	s := memorystore.New()
	dp1 := payment.NewDummyPayment()
	pp1, err := presentation.FromDomainPayment(dp1)
	require.Nil(t, err)
	pp2, err := SavePayment(s, pp1)
	require.Nil(t, err)

	pps, err := ListPayments(s)
	require.Nil(t, err)
	assert.Len(t, pps, 1)
	pp3 := pps[0]
	assert.Equal(t, 1, pp3.Version)
	assert.Equal(t, *pp3, *pp2)
}
