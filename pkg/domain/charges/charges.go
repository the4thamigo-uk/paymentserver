package charges

import (
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/currency"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
)

// Charges holds the information about how fees are distributed.
type Charges struct {
	BearerCode Code
	Sender     map[currency.Currency]money.Money
	Receiver   money.Money
}

// Validate performs some checks on the charges.
// TODO : These business rules need checking
func (c *Charges) Validate(sendCcy currency.Currency, recvCcy currency.Currency) error {
	// NB: we assume that if we are in different currencies we need
	// to have sender charges in both currencies
	// This is just a guess!
	numCharges := len(c.Sender)
	expCharges := 2
	if sendCcy == recvCcy {
		expCharges = 1
	}
	if numCharges != expCharges {
		return errNumberCharges(numCharges)
	}
	_, ok := c.Sender[sendCcy]
	if !ok {
		return errChargeNotFound("sender", sendCcy)
	}
	_, ok = c.Sender[recvCcy]
	if !ok {
		return errChargeNotFound("sender", recvCcy)
	}

	// NB: we assume that the receiver charge is in the currency of the sender.
	// This is just a guess!
	if c.Receiver.Currency() != sendCcy {
		return errChargeNotFound("receiver", recvCcy)
	}
	return nil
}
