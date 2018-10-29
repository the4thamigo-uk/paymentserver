package presentation

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/payment"
	"testing"
)

func TestPresentation_FromPayment(t *testing.T) {
	p1 := payment.NewDummyPayment()
	pp, err := FromDomainPayment(p1)
	require.Nil(t, err)
	p2, err := pp.ToDomainPayment()
	require.Nil(t, err)
	assert.Equal(t, p1, *p2)
}

func TestPresentation_BadID(t *testing.T) {
	p1 := payment.NewDummyPayment()
	pp, err := FromDomainPayment(p1)
	require.Nil(t, err)

	pp.Entity.ID = "INVALID"
	_, err = pp.ToDomainPayment()
	require.NotNil(t, err)
}
