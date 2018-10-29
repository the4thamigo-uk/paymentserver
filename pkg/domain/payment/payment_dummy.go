package payment

import (
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/account"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/amount"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/bank"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/charges"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/date"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/entity"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/fx"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/sponsor"
)

func newDummyBeneficiary() account.Account {
	accType := 0
	return account.Account{
		ID: account.Identifier{
			Name:   "J Smith Ltd",
			Number: "JO71CBJO0000000000001234567890",
			Type:   &accType,
			Code:   account.IBAN},
		Name:    "John Smith",
		Address: "1 Test Street",
		BankID: bank.Identifier{
			ID:   "403000",
			Code: bank.GBDSC},
	}
}

func newDummyDebtor() account.Account {
	return account.Account{
		ID: account.Identifier{
			Name:   "J Doe Ltd",
			Number: "12345678",
			Code:   account.BBAN},
		Name:    "John Doe",
		Address: "2 Test Street",
		BankID: bank.Identifier{
			ID:   "404000",
			Code: bank.GBDSC},
	}
}

func newDummyCharges() charges.Charges {
	return charges.Charges{
		BearerCode: charges.SHAR,
		Sender: []money.Money{
			money.MustParse("234.56", "USD"),
			money.MustParse("123.45", "GBP"),
		},
		Receiver: money.MustParse("234.56", "USD"),
	}
}

func newDummyFx() *fx.Contract {
	return &fx.Contract{
		Reference: "FX123",
		Rate:      amount.MustParse("1.23456789"),
		Domestic:  money.MustParse("2.89", "USD"),
	}
}

func newDummySponsor() sponsor.Sponsor {
	return sponsor.Sponsor{
		Number: "123456789",
		BankID: bank.Identifier{
			ID:   "987654321",
			Code: bank.GBDSC,
		},
	}
}

// NewDummyPayment creates a payment object used for testing purposes only
func NewDummyPayment() Payment {
	return Payment{
		Entity:         entity.MustNew(),
		OrganisationID: "123",
		Credit:         money.MustParse("2.34", "GBP"),
		Beneficiary:    newDummyBeneficiary(),
		Debtor:         newDummyDebtor(),
		Charges:        newDummyCharges(),
		Fx:             newDummyFx(),
		ID:             "123",
		Purpose:        "pay something to someone",
		Type:           "Credit",
		Scheme:         "FPS",
		SchemeType:     "ImmediatePayment",
		SchemeSubType:  "InternetBanking",
		EndToEndRef:    "another ref",
		NumericRef:     "123",
		Reference:      "ref123",
		ProcessingDate: date.MustParse("2000-02-01"),
		Sponsor:        newDummySponsor(),
	}
}
