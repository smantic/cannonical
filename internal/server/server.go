package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Config struct {
	Addr string
	Port string
}

type Server struct {
	addr string
	port string

	mux  *chi.Mux
	prom *PromMiddleware
}

func NewServer(c Config) Server {

	mux := chi.NewRouter()

	mux.Get("/healthz", healthz())
	mux.Handle("/metrics", promhttp.Handler())

	prom := NewPromMiddleware()

	return Server{
		mux:  mux,
		prom: &prom,
	}
}

func (s *Server) ApplyRoutes() {
	//routes := map[string]http.HandlerFunc{
	//	"/v1/endpoint": handlers(healthz(), s.prom.CountRequest)
	//}

	//for route, handler := range routes {
	//	s.mux.Handle(route, handler)
	//}
}

// Run will run our http server
func (s *Server) Run() error {
	return http.ListenAndServe(s.addr+":"+s.port, s.mux)
}
