package tests

import (
	"github.com/the4thamigo-uk/paymentserver/pkg/presentation"
)

func newDummyPayment() *presentation.Payment {
	accType := 0
	return &presentation.Payment{
		Type: "Payment",
		Entity: presentation.Entity{
			ID:      "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
			Version: 0,
		},
		OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
		Attributes: presentation.Attributes{
			Amount: "100.21",
			BeneficiaryParty: presentation.Account{
				AccountName:       "W Owens",
				AccountNumber:     "31926819",
				AccountNumberCode: "BBAN",
				AccountType:       &accType,
				Address:           "1 The Beneficiary Localtown SE2",
				BankID:            "403000",
				BankIDCode:        "GBDSC",
				Name:              "Wilfred Jeremiah Owens",
			},
			ChargesInformation: presentation.Charges{
				BearerCode: "SHAR",
				SenderCharges: []presentation.Money{
					presentation.Money{
						Amount:   "5.00",
						Currency: "GBP",
					},
					presentation.Money{
						Amount:   "10.00",
						Currency: "USD",
					},
				},
				ReceiverChargesAmount:   "1.00",
				ReceiverChargesCurrency: "USD",
			},
			Currency: "GBP",
			DebtorParty: presentation.Account{
				AccountName:       "EJ Brown Black",
				AccountNumber:     "BG18RZBB91550123456789",
				AccountNumberCode: "IBAN",
				Address:           "10 Debtor Crescent Sourcetown NE1",
				BankID:            "203301",
				BankIDCode:        "GBDSC",
				Name:              "Emelia Jane Brown",
			},
			EndToEndReference: "Wil piano Jan",
			Fx: &presentation.Fx{
				ContractReference: "FX123",
				ExchangeRate:      "2.00000",
				OriginalAmount:    "200.42",
				OriginalCurrency:  "USD",
			},
			NumericReference:     "1002001",
			PaymentID:            "123456789012345678",
			PaymentPurpose:       "Paying for goods/services",
			PaymentScheme:        "FPS",
			PaymentType:          "Credit",
			ProcessingDate:       "2017-01-18",
			Reference:            "Payment for Em's piano lessons",
			SchemePaymentSubType: "InternetBanking",
			SchemePaymentType:    "ImmediatePayment",
			SponsorParty: presentation.Sponsor{
				AccountNumber: "56781234",
				BankID:        "123123",
				BankIDCode:    "GBDSC",
			},
		},
	}
}
