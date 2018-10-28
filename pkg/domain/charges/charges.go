package charges

import (
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/currency"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
)

// Charges holds the information about how fees are distributed.
type Charges struct {
	BearerCode Code          `json:"bearer_code"`
	Sender     []money.Money `json:"sender"`
	Receiver   money.Money   `json:"receiver"`
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
	_, i := c.senderCharge(sendCcy)
	if i < 0 {
		return errChargeNotFound("sender", sendCcy)
	}
	_, i = c.senderCharge(recvCcy)
	if i < 0 {
		return errChargeNotFound("sender", recvCcy)
	}

	// NB: we assume that the receiver charge is in the currency of the sender.
	// This is just a guess!
	if c.Receiver.Currency() != sendCcy {
		return errChargeNotFound("receiver", recvCcy)
	}
	return nil
}

func (c *Charges) senderCharge(ccy currency.Currency) (money.Money, int) {
	for i, m := range c.Sender {
		if m.Currency() == ccy {
			return m, i
		}
	}
	return money.Money{}, -1
}
