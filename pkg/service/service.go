package service

import (
	"github.com/imdario/mergo"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/payment"
	"github.com/the4thamigo-uk/paymentserver/pkg/presentation"
	"github.com/the4thamigo-uk/paymentserver/pkg/store"
)

// LoadPayment loads the latest payment from the store
func LoadPayment(s store.Store, id string) (*presentation.Payment, error) {
	var dp payment.Payment
	_, err := s.Load(store.NewID(id, 0), &dp)
	if err != nil {
		return nil, err
	}
	dp.Entity.Version = 1
	return presentation.FromDomainPayment(dp)
}

// SavePayment writes the complete payment to the store
func SavePayment(s store.Store, pp *presentation.Payment) (*presentation.Payment, error) {
	dp, err := pp.ToDomainPayment()
	if err != nil {
		return nil, err
	}
	sid := store.NewID(dp.Entity.ID.String(), dp.Entity.Version)
	sid, err = s.Save(sid, &dp)
	if err != nil {
		return nil, err
	}
	dp.Entity.Version = sid.Version
	return presentation.FromDomainPayment(*dp)
}

// UpdatePayment patches the attributes of the payment with the attributes in the payment provided
func UpdatePayment(s store.Store, pp *presentation.Payment) (*presentation.Payment, error) {
	spp, err := LoadPayment(s, pp.ID.String())
	if err != nil {
		return nil, err
	}
	err = mergo.Merge(&pp.Attributes, spp.Attributes)
	if err != nil {
		return nil, err
	}
	return SavePayment(s, pp)
}
