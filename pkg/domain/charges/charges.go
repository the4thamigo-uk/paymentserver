package charges

import (
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
)

type Charges struct {
	BearerCode      Code
	SenderCharges   []money.Money
	ReceiverCharges money.Money
}
