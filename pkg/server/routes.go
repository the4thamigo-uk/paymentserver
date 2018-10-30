package server

import (
	"net/http"
	"regexp"
)

type route struct {
	path        string
	uriTmpl     string
	method      string
	handler     handler
	title       string
	linksRels   []string
	linksRoutes routes
}

type routes map[string]*route

var uriTmplRegex = regexp.MustCompile(`:([^/]*)`)

const (
	relIndex         = "index"
	relPaymentList   = "payment:list"
	relPaymentCreate = "payment:create"
	relPaymentLoad   = "payment:load"
	relPaymentSave   = "payment:save"
	relPaymentUpdate = "payment:update"
	relPaymentDelete = "payment:delete"
)

func newRoutes() routes {
	listRels := []string{
		relIndex,
		relPaymentCreate,
		relPaymentList,
	}
	paymentRels := []string{
		relIndex,
		relPaymentList,
		relPaymentCreate,
		relPaymentLoad,
		relPaymentSave,
		relPaymentUpdate,
		relPaymentDelete,
	}
	rs := routes{
		relIndex: &route{
			title:     "Payments server",
			path:      "/",
			method:    http.MethodGet,
			handler:   rootIndex,
			linksRels: listRels,
		},
		relPaymentList: &route{
			title:     "List payments",
			path:      "/payments",
			method:    http.MethodGet,
			handler:   paymentList,
			linksRels: listRels,
		},
		relPaymentCreate: &route{
			title:     "Create payment",
			path:      "/payments",
			method:    http.MethodPost,
			handler:   paymentCreate,
			linksRels: paymentRels,
		},
		relPaymentLoad: &route{
			title:     "Load payment",
			path:      "/payments/:id/:version",
			method:    http.MethodGet,
			handler:   paymentLoad,
			linksRels: paymentRels,
		},
		relPaymentSave: &route{
			title:     "Save a payment",
			path:      "/payments/:id/:version",
			method:    http.MethodPut,
			handler:   paymentSave,
			linksRels: paymentRels,
		},
		relPaymentUpdate: &route{
			title:     "Update a payment",
			path:      "/payments/:id/:version",
			method:    http.MethodPatch,
			handler:   paymentUpdate,
			linksRels: paymentRels,
		},
		relPaymentDelete: &route{
			title:     "Delete a payment",
			path:      "/payments/:id/:version",
			method:    http.MethodDelete,
			handler:   paymentDelete,
			linksRels: paymentRels,
		},
	}
	for _, r := range rs {
		r.init(rs)
	}
	return rs
}

func (r *route) init(rs routes) {
	r.uriTmpl = uriTmplRegex.ReplaceAllString(r.path, `{$1}`)

	lrs := routes{}
	for _, lrel := range r.linksRels {
		lr, ok := rs[lrel]
		if !ok {
			continue
		}
		lrs[lrel] = lr
	}
	r.linksRoutes = lrs
}
