package server

import (
	"github.com/julienschmidt/httprouter"
	"github.com/the4thamigo-uk/paymentserver/pkg/store"
	"net/http"
)

type handler func(r *request, g *globals) (*response, error)

type request struct {
	r  *http.Request
	p  httprouter.Params
	w  http.ResponseWriter
	rt *route
}

type errorData struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

type response struct {
	Error *errorData  `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Links links       `json:"_links,omitempty"`
}

type globals struct {
	cfg    *Config
	routes routes
	store  store.Store
}

func rootIndex(r *request, g *globals) (*response, error) {
	return &response{
		Links: linksForRoute(r.rt, nil),
	}, nil
}
