package server

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/the4thamigo-uk/paymentserver/pkg/store"
	"github.com/the4thamigo-uk/paymentserver/pkg/store/memorystore"
	"net/http"
	"time"
)

// Server is an instance of a REST paymentserver.
type Server struct {
	s *http.Server
	g *globals
}

// NewServer creates a new instance of Server with default in-memory storage,
func NewServer(cfg *Config) *Server {
	g := &globals{
		cfg:    cfg,
		routes: newRoutes(),
		store:  memorystore.New(),
	}

	r := httprouter.New()
	for _, route := range g.routes {
		r.Handle(route.method, route.path, rootHandler(route, g))
	}
	return &Server{
		s: &http.Server{
			Addr:    cfg.Address,
			Handler: r,
		},
		g: g,
	}
}

// WithStore overrides the default in-memory storage with a client-specified storage.
func (s *Server) WithStore(store store.Store) *Server {
	s.g.store = store
	return s
}

// ListenAndServe starts the server. This function blocks until either Shutdown is called or the process is terminated.
func (s *Server) ListenAndServe() error {
	return s.s.ListenAndServe()
}

// Shutdown attempts a graceful shutdown of the server, timing out after 5 seconds.
func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.s.Shutdown(ctx)
}

func rootHandler(rt *route, g *globals) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := &request{
			r:  r,
			w:  w,
			p:  p,
			rt: rt,
		}
		rsp, err := rt.handler(ctx, g)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			rsp = errResponse(err, http.StatusBadRequest)
		}
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/hal+json")
		_, err = w.Write(b)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func errResponse(err error, code int) *response {
	return &response{
		Error: &errorData{
			Message: err.Error(),
			Code:    code,
		}}
}
