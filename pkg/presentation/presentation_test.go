package presentation

import (
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/payment"
	"testing"
)

func TestPresentation_FromPayment(t *testing.T) {
	p1 := payment.NewDummyPayment()
	pID := uuid.Must(uuid.NewV4())
	orgID := uuid.Must(uuid.NewV4())
	pp, err := FromDomainPayment(pID, 1, orgID, p1)
	require.Nil(t, err)
	p2, err := pp.ToDomainPayment()
	require.Nil(t, err)
	assert.Equal(t, p1, *p2)
}
