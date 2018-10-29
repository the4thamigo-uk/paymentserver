package server

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/the4thamigo-uk/paymentserver/pkg/presentation"
	"github.com/the4thamigo-uk/paymentserver/pkg/service"
	"io"
	"io/ioutil"
	"strconv"
)

func paymentCreate(r *request, g *globals) (*response, error) {
	var pp presentation.Payment
	err := unmarshalReader(r.r.Body, &pp)
	if err != nil {
		return nil, err
	}
	pp2, err := service.CreatePayment(g.store, &pp)
	if err != nil {
		return nil, err
	}
	pps := []*presentation.Payment{pp2}
	return paymentResponse(r.rt, &pp2.Entity, pps), nil
}

func paymentSave(r *request, g *globals) (*response, error) {
	eid, err := paymentID(r)
	if err != nil {
		return nil, err
	}
	var pp presentation.Payment
	err = unmarshalReader(r.r.Body, &pp)
	if err != nil {
		return nil, err
	}
	pp.Entity = eid
	pp2, err := service.SavePayment(g.store, &pp)
	if err != nil {
		return nil, err
	}
	pps := []*presentation.Payment{pp2}
	return paymentResponse(r.rt, &pp2.Entity, pps), nil
}

func paymentUpdate(r *request, g *globals) (*response, error) {
	eid, err := paymentID(r)
	if err != nil {
		return nil, err
	}
	var pp presentation.Payment
	err = unmarshalReader(r.r.Body, &pp)
	if err != nil {
		return nil, err
	}
	pp.Entity = eid
	pp2, err := service.UpdatePayment(g.store, &pp)
	if err != nil {
		return nil, err
	}
	pps := []*presentation.Payment{pp2}
	return paymentResponse(r.rt, &pp2.Entity, pps), nil
}

func paymentLoad(r *request, g *globals) (*response, error) {
	eid, err := paymentID(r)
	if err != nil {
		return nil, err
	}
	pp, err := service.LoadPayment(g.store, eid)
	if err != nil {
		return nil, err
	}
	pps := []*presentation.Payment{pp}
	return paymentResponse(r.rt, &pp.Entity, pps), nil
}

func paymentDelete(r *request, g *globals) (*response, error) {
	eid, err := paymentID(r)
	if err != nil {
		return nil, err
	}
	pp, err := service.DeletePayment(g.store, eid)
	if err != nil {
		return nil, err
	}
	pps := []*presentation.Payment{pp}
	return paymentResponse(r.rt, &pp.Entity, pps), nil
}

func paymentList(r *request, g *globals) (*response, error) {
	pps, err := service.ListPayments(g.store)
	if err != nil {
		return nil, err
	}
	return paymentResponse(r.rt, nil, pps), nil
}

func paymentID(r *request) (presentation.Entity, error) {
	id := r.p.ByName("id")

	ver := 0
	if s := r.p.ByName("version"); s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			return presentation.Entity{}, errors.Wrap(err, "Version must be an integer")
		}
		ver = i
	}
	return presentation.NewEntity(id, ver), nil
}

func paymentResponse(rt *route, eid *presentation.Entity, pps []*presentation.Payment) *response {
	vals := map[string]interface{}{}
	if eid != nil {
		vals["id"] = eid.ID
		vals["version"] = eid.Version
	}
	return &response{
		Data:  pps,
		Links: linksForRoute(rt, vals),
	}
}

func unmarshalReader(r io.Reader, obj interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, obj)
}
