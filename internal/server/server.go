package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

// Config is the config for the server
type Config struct {
	Addr string
	Port string

	Logger *zap.Logger
}

// Server is our http server
type Server struct {
	addr string
	port string

	mux *chi.Mux
}

// Handlers contains all of our handlers.
type Handlers struct {
	Prom *PromMiddleware
	Log  *ZapMiddleware
}

// NewServer will create a new server
func NewServer(c Config) Server {

	mux := chi.NewRouter()

	prom := NewPromMiddleware()
	log := NewZapMiddleware(c.Logger)

	h := Handlers{
		Prom: &prom,
		Log:  &log,
	}

	h.applyRoutes(mux)

	return Server{
		mux: mux,
	}
}

func (h *Handlers) applyRoutes(mux *chi.Mux) {

	mux.Use(middleware.Recoverer)
	mux.Use(h.Prom.RecordRequest)
	mux.Use(h.Log.LogRequest)

	mux.Get("/healthz", healthz())
	mux.Handle("/metrics", promhttp.Handler())

	// mounts pprof
	mux.Mount("/debug", middleware.Profiler())
}

// Run will run our http server
func (s *Server) Run() error {
	return http.ListenAndServe(s.addr+":"+s.port, s.mux)
}
