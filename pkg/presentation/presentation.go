package presentation

import (
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/account"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/amount"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/bank"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/charges"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/date"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/entity"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/fx"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/money"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/payment"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/sponsor"
)

// Payment is an external representation of the associated domain object.
type Payment struct {
	Type           string     `json:"type"`
	ID             entity.ID  `json:"id"`
	Version        int        `json:"version"`
	OrganisationID string     `json:"organisation_id"`
	Attributes     Attributes `json:"attributes"`
}

// Attributes is an external representation of the associated domain object.
type Attributes struct {
	Amount               string  `json:"amount"`
	BeneficiaryParty     Account `json:"beneficiary_party"`
	DebtorParty          Account `json:"debtor_party"`
	ChargesInformation   Charges `json:"charges_information"`
	Currency             string  `json:"currency"`
	EndToEndReference    string  `json:"end_to_end_reference"`
	Fx                   *Fx     `json:"fx,omitempty"`
	NumericReference     string  `json:"numeric_reference"`
	PaymentID            string  `json:"payment_id"`
	PaymentPurpose       string  `json:"payment_purpose"`
	PaymentScheme        string  `json:"payment_scheme"`
	PaymentType          string  `json:"payment_type"`
	ProcessingDate       string  `json:"processing_date"`
	Reference            string  `json:"reference"`
	SchemePaymentSubType string  `json:"scheme_payment_sub_type"`
	SchemePaymentType    string  `json:"scheme_payment_type"`
	SponsorParty         Sponsor `json:"sponsor_party"`
}

// Account is an external representation of the associated domain object.
type Account struct {
	AccountName       string `json:"account_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	AccountType       *int   `json:"account_type"`
	Address           string `json:"address"`
	BankID            string `json:"bank_id"`
	BankIDCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

// Charges is an external representation of the associated domain object object.
type Charges struct {
	BearerCode              string  `json:"bearer_code"`
	SenderCharges           []Money `json:"sender_charges"`
	ReceiverChargesAmount   string  `json:"receiver_charges_amount"`
	ReceiverChargesCurrency string  `json:"receiver_charges_currency"`
}

// Fx is an external representation of the associated domain object object.
type Fx struct {
	ContractReference string `json:"contract_reference"`
	ExchangeRate      string `json:"exchange_rate"`
	OriginalAmount    string `json:"original_amount"`
	OriginalCurrency  string `json:"original_currency"`
}

// Money is an external representation of the associated domain object object.
type Money struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// Sponsor is an external representation of the associated domain object object.
type Sponsor struct {
	AccountNumber string `json:"account_number"`
	BankID        string `json:"bank_id"`
	BankIDCode    string `json:"bank_id_code"`
}

// FromDomainPayment creates an external representation of the domain object
func FromDomainPayment(p payment.Payment) (*Payment, error) {
	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return &Payment{
		Type:           "Payment",
		ID:             p.Entity.ID,
		Version:        p.Entity.Version,
		OrganisationID: p.OrganisationID,
		Attributes: Attributes{
			Amount:               p.Credit.Amount().String(),
			Currency:             p.Credit.Currency().String(),
			BeneficiaryParty:     FromDomainAccount(p.Beneficiary),
			DebtorParty:          FromDomainAccount(p.Debtor),
			ChargesInformation:   FromDomainCharges(p.Charges),
			EndToEndReference:    p.EndToEndRef,
			Fx:                   FromDomainFx(p.Fx),
			NumericReference:     p.NumericRef,
			PaymentID:            p.ID,
			PaymentPurpose:       p.Purpose,
			PaymentScheme:        p.Scheme,
			PaymentType:          p.Type,
			ProcessingDate:       p.ProcessingDate.String(),
			Reference:            p.Reference,
			SchemePaymentSubType: p.SchemeSubType,
			SchemePaymentType:    p.SchemeType,
			SponsorParty:         FromDomainSponsor(p.Sponsor),
		}}, nil
}

// FromDomainAccount creates an external representation of the domain object
func FromDomainAccount(a account.Account) Account {
	return Account{
		AccountName:       a.ID.Name,
		AccountNumber:     a.ID.Number,
		AccountNumberCode: a.ID.Code.String(),
		AccountType:       a.ID.Type,
		Address:           a.Address,
		BankID:            a.BankID.ID,
		BankIDCode:        a.BankID.Code.String(),
		Name:              a.Name,
	}
}

// FromDomainFx creates an external representation of the domain object
func FromDomainFx(fx *fx.Contract) *Fx {
	if fx == nil {
		return nil
	}
	return &Fx{
		ContractReference: fx.Reference,
		ExchangeRate:      fx.Rate.String(),
		OriginalAmount:    fx.Domestic.Amount().String(),
		OriginalCurrency:  fx.Domestic.Currency().String(),
	}
}

// FromDomainCharges creates an external representation of the domain object
func FromDomainCharges(c charges.Charges) Charges {
	sc := []Money{}
	for _, s := range c.Sender {
		sc = append(sc, Money{
			Amount:   s.Amount().String(),
			Currency: s.Currency().String(),
		})
	}
	return Charges{
		BearerCode:              c.BearerCode.String(),
		SenderCharges:           sc,
		ReceiverChargesAmount:   c.Receiver.Amount().String(),
		ReceiverChargesCurrency: c.Receiver.Currency().String(),
	}
}

// FromDomainSponsor creates an external representation of the domain object
func FromDomainSponsor(s sponsor.Sponsor) Sponsor {
	return Sponsor{
		AccountNumber: s.Number,
		BankID:        s.BankID.ID,
		BankIDCode:    s.BankID.Code.String(),
	}
}

// ToDomainPayment creates a domain object from the associated external representation
func (p Payment) ToDomainPayment() (*payment.Payment, error) {
	a := &p.Attributes

	credit, err := money.Parse(a.Amount, a.Currency)
	if err != nil {
		return nil, err
	}
	beneficiary, err := a.BeneficiaryParty.ToDomainAccount()
	if err != nil {
		return nil, err
	}
	debtor, err := a.DebtorParty.ToDomainAccount()
	if err != nil {
		return nil, err
	}
	charges, err := a.ChargesInformation.ToDomainCharges()
	if err != nil {
		return nil, err
	}
	fx, err := a.Fx.ToDomainFx()
	if err != nil {
		return nil, err
	}
	sponsor, err := a.SponsorParty.ToDomainSponsor()
	if err != nil {
		return nil, err
	}
	p2 := &payment.Payment{
		Entity: entity.Entity{
			ID:      p.ID,
			Version: p.Version,
		},
		OrganisationID: p.OrganisationID,
		Credit:         *credit,
		Beneficiary:    *beneficiary,
		Debtor:         *debtor,
		ProcessingDate: date.MustParse("2000-02-01"),
		Charges:        *charges,
		Fx:             fx,
		Sponsor:        *sponsor,
		ID:             a.PaymentID,
		Type:           a.PaymentType,
		Purpose:        a.PaymentPurpose,
		Scheme:         a.PaymentScheme,
		SchemeType:     a.SchemePaymentType,
		SchemeSubType:  a.SchemePaymentSubType,
		NumericRef:     a.NumericReference,
		EndToEndRef:    a.EndToEndReference,
		Reference:      a.Reference,
	}
	err = p2.Validate()
	return p2, err
}

// ToDomainAccount creates a domain object from the associated external representation
func (a Account) ToDomainAccount() (*account.Account, error) {
	acode, err := account.Parse(a.AccountNumberCode)
	if err != nil {
		return nil, err
	}
	bcode, err := bank.Parse(a.BankIDCode)
	if err != nil {
		return nil, err
	}
	return &account.Account{
		ID: account.Identifier{
			Name:   a.AccountName,
			Number: a.AccountNumber,
			Code:   *acode,
			Type:   a.AccountType,
		},
		Name:    a.Name,
		Address: a.Address,
		BankID: bank.Identifier{
			ID:   a.BankID,
			Code: *bcode,
		},
	}, nil
}

// ToDomainCharges creates a domain object from the associated external representation
func (c *Charges) ToDomainCharges() (*charges.Charges, error) {
	sc := []money.Money{}
	for _, s := range c.SenderCharges {
		m, err := money.Parse(s.Amount, s.Currency)
		if err != nil {
			return nil, err
		}
		sc = append(sc, *m)
	}
	rc, err := money.Parse(c.ReceiverChargesAmount, c.ReceiverChargesCurrency)
	if err != nil {
		return nil, err
	}
	bc, err := charges.Parse(c.BearerCode)
	if err != nil {
		return nil, err
	}
	return &charges.Charges{
		BearerCode: *bc,
		Sender:     sc,
		Receiver:   *rc,
	}, nil
}

// ToDomainFx creates a domain object from the associated external representation
func (f *Fx) ToDomainFx() (*fx.Contract, error) {
	r, err := amount.Parse(f.ExchangeRate)
	if err != nil {
		return nil, err
	}
	d, err := money.Parse(f.OriginalAmount, f.OriginalCurrency)
	if err != nil {
		return nil, err
	}
	return &fx.Contract{
		Reference: f.ContractReference,
		Rate:      r,
		Domestic:  *d,
	}, nil
}

// ToDomainSponsor creates a domain object from the associated external representation
func (s *Sponsor) ToDomainSponsor() (*sponsor.Sponsor, error) {
	bc, err := bank.Parse(s.BankIDCode)
	if err != nil {
		return nil, err
	}
	return &sponsor.Sponsor{
		Number: s.AccountNumber,
		BankID: bank.Identifier{
			ID:   s.BankID,
			Code: *bc,
		},
	}, nil
}
