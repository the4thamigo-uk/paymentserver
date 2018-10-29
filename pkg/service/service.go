package service

import (
	"github.com/imdario/mergo"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/entity"
	"github.com/the4thamigo-uk/paymentserver/pkg/domain/payment"
	"github.com/the4thamigo-uk/paymentserver/pkg/presentation"
	"github.com/the4thamigo-uk/paymentserver/pkg/store"
)

// CreatePayment creates a new payment and saves it to the store
// The payment must be valid
// The ID and Version will be overwritten.
func CreatePayment(s store.Store, pp *presentation.Payment) (*presentation.Payment, error) {
	eid, err := entity.New()
	if err != nil {
		return nil, err
	}
	pp.ID = eid.ID
	pp.Version = eid.Version
	return SavePayment(s, pp)
}

// LoadPayment loads the version of the payment from the store.
// A version of 0 retrieves the latest version
func LoadPayment(s store.Store, eid presentation.Entity) (*presentation.Payment, error) {
	var dp payment.Payment
	sid := toStoreID(eid)
	sid, err := s.Load(sid, &dp)
	if err != nil {
		return nil, err
	}
	dp.Entity.Version = sid.Version
	return presentation.FromDomainPayment(dp)
}

// SavePayment writes the complete payment to the store
func SavePayment(s store.Store, pp *presentation.Payment) (*presentation.Payment, error) {
	dp, err := pp.ToDomainPayment()
	if err != nil {
		return nil, err
	}
	sid := toStoreID(pp.Entity)
	sid, err = s.Save(sid, &dp)
	if err != nil {
		return nil, err
	}
	dp.Entity.Version = sid.Version
	return presentation.FromDomainPayment(*dp)
}

// UpdatePayment patches the attributes of the payment with the attributes in the payment provided
func UpdatePayment(s store.Store, pp *presentation.Payment) (*presentation.Payment, error) {
	spp, err := LoadPayment(s, pp.Entity)
	if err != nil {
		return nil, err
	}
	err = mergo.Merge(&pp.Attributes, spp.Attributes)
	if err != nil {
		return nil, err
	}
	return SavePayment(s, pp)
}

// DeletePayment removes the payment from the store and returns the removed payment
func DeletePayment(s store.Store, eid presentation.Entity) (*presentation.Payment, error) {
	var dp payment.Payment
	sid := toStoreID(eid)
	sid, err := s.Delete(sid, &dp)
	if err != nil {
		return nil, err
	}
	dp.Entity.Version = sid.Version
	return presentation.FromDomainPayment(dp)
}

func toStoreID(eid presentation.Entity) store.ID {
	return store.NewID(eid.ID, eid.Version)
}
